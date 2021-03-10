package sched

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"container/heap"
	"context"
	"runtime"
	"sync/atomic"
	"unsafe"
)

type Future [R any] struct {
	Err error
	value atomic.Value
}

// Get implements Future interface
func (f *Future[R]) Get() R {
	var v interface{}
	// spin until value is stored in future.value
	for ; v == nil && f.Err == nil; v = f.value.Load() {
		runtime.Gosched()
	}
	return v.(R)
}

func (f *Future[R]) put(v R) {
	f.value.Store(v)
}


// Task interface for sched
type Task [R any] interface {
	// GetID must returns a unique ID for all of the scheduled task.
	GetID() (id string)
	// GetExecution returns the time for task execution.
	GetExecution() (execute time.Time)
	// GetRetryTime returns the retry time if a task was failed.
	GetRetryTime() (execute time.Time)
	// Execute executes the actual task, it can return a result,
	// or if the task need a retry, or it was failed in this execution.
	Execute() (result R, retry bool, fail error)
}



// sched is the actual scheduler for task scheduling
//
// sched implements greedy scheduling, with a timer and a task queue,
// the task queue is a priority queue that orders tasks by executing
// time. The timer is the only time.Timer lives in runtime, it serves
// the head task in the task queue.
type sched [R any, T Task[R]] struct {
	// running counts the tasks already starts that cannot be stopped.
	running uint64 // atomic
	// pausing is a sign that indicates if sched should stop running.
	pausing uint64 // atomic
	// timer is the only timer during the runtime
	timer unsafe.Pointer // *time.Timer
	// cancel cancels a timer if a timer need to reset
	cancel atomic.Value // context.CancelFunc
	// tasks is a TaskQueue that stores all unscheduled tasks in memory
	tasks *taskQueue[R]
}

// NewSched returns a scheduler that schedules type T tasks.
func NewSched[R any, T Task[R]]() *sched[R, T] {
	return &sched[R, T]{
		timer: unsafe.Pointer(time.NewTimer(0)),
		tasks: newTaskQueue[R](),
	}
}

// Stop stops runtime scheduler gracefully.
// Note that the call should only be called then application terminates
func (sched0 *sched[R, T]) Stop() {
	sched0.Pause()

	// wait until all started tasks
	for atomic.LoadUint64(&sched0.running) > 0 {
		runtime.Gosched()
	}

	// reset pausing indicator
	atomic.AddUint64(&sched0.pausing, ^uint64(0))
}

// Wait waits all tasks to be scheduled.
func (sched0 *sched[R, T]) Wait() {
	// With function call
	for sched0.tasks.length() != 0 {
		runtime.Gosched()
	}
}

// Submit given tasks
func (sched0 *sched[R, T]) Submit(t T) *Future[R] {
	return sched0.schedule(t, t.GetExecution())
}

// Trigger given tasks immediately
func (sched0 *sched[R, T]) Trigger(t T) *Future[R] {
	return sched0.schedule(t, time.Now())
}

// Pause stops the sched timing
func (sched0 *sched[R, T]) Pause() {
	atomic.AddUint64(&sched0.pausing, 1)
	sched0.pause()
}

// Resume resumes sched and start executing tasks
// this is a pair call with Pause(), Resume() must be called second
func (sched0 *sched[R, T]) Resume() {
	atomic.AddUint64(&sched0.pausing, ^uint64(0)) // -1
	sched0.resume()
}


func (s *sched[R, T]) schedule(t T, when time.Time) *Future[R] {
	s.pause()

	// if priority is able to be update
	if future, ok := s.tasks.update(t, when); ok {
		s.resume()
		return future
	}

	future := s.tasks.push(newTaskItem[R](t, when))
	s.resume()
	return future
}

func (s *sched[R, T]) reschedule(t *task[R], when time.Time) {
	s.pause()
	t.priority = when
	s.tasks.push(t)
	s.resume()
}

func (s *sched[R, T]) getTimer() (t *time.Timer) {
	for {
		t = (*time.Timer)(atomic.LoadPointer(&s.timer))
		if t != nil {
			return
		}
		runtime.Gosched()
	}
}

func (s *sched[R, T]) setTimer(d time.Duration) {
	for {
		// fast path: reuse the timer
		old := atomic.SwapPointer(&s.timer, nil)
		if old != nil {
			if (*time.Timer)(old).Stop() {
				(*time.Timer)(old).Reset(d)
				if atomic.CompareAndSwapPointer(&s.timer, nil, old) {
					return
				}
				runtime.Gosched()
				continue
			}
		}

		// slow path: fail to stop, use a new timer.
		// this happens only if the sched is super busy.
		if atomic.CompareAndSwapPointer(&s.timer, old,
			unsafe.Pointer(time.NewTimer(d))) {
			if old != nil {
				(*time.Timer)(old).Stop()
			}
			return
		}
		runtime.Gosched()
	}
}

// pause pauses sched without pause tasks from running
func (s *sched[R, T]) pause() {
	old := atomic.LoadPointer(&s.timer)
	// if old is nil then there is someone who tries to stop the timer.
	if old != nil {
		(*time.Timer)(old).Stop()
	}
}

func (s *sched[R, T]) resume() {
	// Cancel as soon as possible, this must happens before setTimer
	ctx, cancel := context.WithCancel(context.Background())
	if x, ok := s.cancel.Load().(context.CancelFunc); ok {
		x()
	}
	s.cancel.Store(cancel)

	t := s.tasks.peek()
	if t == nil {
		return
	}
	s.setTimer(t.GetExecution().Sub(time.Now()))

	go func(ctx context.Context) {
		select {
		case <-s.getTimer().C:
			s.worker()
		case <-ctx.Done():
		}
	}(ctx)
}

func (s *sched[R, T]) worker() {
	// fast path.
	// if sched requires pausing, then stop executing and resume it.
	if atomic.LoadUint64(&s.pausing) > 0 {
		return
	}

	// medium path.
	// stop execution if task queue is empty
	task := s.tasks.pop()
	if task == nil {
		return
	}

	s.resume()
	s.arrival(task)
}

func (s *sched[R, T]) arrival(t *task[R]) {
	// record running tasks
	atomic.AddUint64(&s.running, 1)
	s.execute(t)
	atomic.AddUint64(&s.running, ^uint64(0)) // -1
}

func (s *sched[R, T]) execute(t *task[R]) {
	defer func() {
		if r := recover(); r != nil {
			t.future.Err = fmt.Errorf(
				"sched: task %s panic while executing, reason: %v",
				t.value.GetID(), r)
		}
	}()

	// for timer tollerance
	if t.value.GetExecution().After(time.Now()) {
		// reschedule task, we must save the task again by using s.Setup
		s.reschedule(t, t.value.GetExecution())
		return
	}
	result, retry, err := t.value.Execute()
	if retry || err != nil {
		s.reschedule(t, t.value.GetRetryTime())
		return
	}
	t.future.put(result)
}

// TaskQueue implements a timer queue based on a heap
// Its supports bi-direction accessing, such as access value by key
// or access key by its value
//
// TODO: lock-free
type taskQueue [R any] struct {
	heap   *taskHeap[R]
	lookup map[string]*task[R]
	mu     sync.Mutex
}

func newTaskQueue [R any] () *taskQueue[R] {
	return &taskQueue[R]{
		heap:   &taskHeap[R]{},
		lookup: map[string]*task[R]{},
	}
}

// length of queue
func (m *taskQueue[R]) length() (l int) {
	m.mu.Lock()
	l = m.heap.Len()
	m.mu.Unlock()
	return
}

// push item
func (m *taskQueue[R]) push(t *task[R]) *Future[R] {
	m.mu.Lock()
	heap.Push(m.heap, t)          // O(log(n))
	m.lookup[t.value.GetID()] = t // O(1)
	m.mu.Unlock()
	return t.future
}

// Pop item
func (m *taskQueue[R]) pop() *task[R] {
	m.mu.Lock()

	if m.heap.Len() == 0 {
		m.mu.Unlock()
		return nil
	}

	item := heap.Pop(m.heap).(*task[R])     // O(log(n))
	delete(m.lookup, item.value.GetID()) // O(1) amortized
	m.mu.Unlock()
	return item
}

// peek the top priority item without deletion
func (m *taskQueue[R]) peek() (t Task[R]) {
	m.mu.Lock()

	if m.heap.Len() == 0 {
		m.mu.Unlock()
		return nil
	}
	t = (*m.heap)[0].value
	m.mu.Unlock()
	return
}

// update of a given task
func (m *taskQueue[R]) update(t Task[R], when time.Time) (*Future[R], bool) {
	m.mu.Lock()
	item, ok := m.lookup[t.GetID()]
	if !ok {
		m.mu.Unlock()
		return nil, false
	}

	item.priority = when
	item.value = t
	heap.Fix(m.heap, item.index) // O(log(n))
	m.mu.Unlock()
	return item.future, true
}

// a task is something we manage in a priority queue.
type task [R any] struct {
	value Task[R] // for storage

	// The index is needed by update and is maintained by the
	// heap.Interface methods.
	index    int       // The index of the item in the heap.
	priority time.Time // type of time for priority
	future   *Future[R]
}

// NewTaskItem creates a new queue item
func newTaskItem[R any](t Task[R], when time.Time) *task[R] {
	return &task[R]{value: t, priority: when, future: &Future[R]{}}
}

type taskHeap[R any] []*task[R]

func (pq taskHeap[R]) Len() int {
	return len(pq)
}

func (pq taskHeap[R]) Less(i, j int) bool {
	return pq[i].priority.Before(pq[j].priority)
}

func (pq taskHeap[R]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *taskHeap[R]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *taskHeap[R]) Push(x interface{}) {
	item := x.(*task[R])
	item.index = len(*pq)
	*pq = append(*pq, item)
}

// =====================================================================


// CustomTask implements task.Interface
type CustomTask struct {
	Public    string
	id        string
	execution time.Time
}

// NewCustomTask creates a task
func NewCustomTask(id string, e time.Time) *CustomTask {
	return &CustomTask{
		Public:    "not nil",
		id:        id,
		execution: e,
	}
}

// GetID get task id
func (t *CustomTask) GetID() (id string) {
	id = t.id
	return
}

// GetExecution get execution time
func (t *CustomTask) GetExecution() (execute time.Time) {
	execute = t.execution
	return
}

// GetTimeout get timeout of execution
func (t *CustomTask) GetTimeout() (executeTimeout time.Duration) {
	return time.Second
}

// GetRetryTime get retry execution duration
func (t *CustomTask) GetRetryTime() time.Time {
	return time.Now().UTC().Add(time.Second)
}

// SetID sets the id of a task
func (t *CustomTask) SetID(id string) {
	t.id = id
}

// IsValidID check id is valid
func (t *CustomTask) IsValidID() bool {
	return true
}

// SetExecution sets the execution time of a task
func (t *CustomTask) SetExecution(current time.Time) (old time.Time) {
	old = t.execution
	t.execution = current
	return
}

// Execute is the actual execution block
func (t *CustomTask) Execute() (r string, retry bool, fail error) {
	O.Push(t.id)
	return fmt.Sprintf("execute task %s.", t.id), false, nil
}


func TestSchedMasiveSchedule(t *testing.T) {
	O.Clear()
	nTasks := 100
	sched0 := NewSched[string, *CustomTask]()
	futures := make([]*Future[string], nTasks)

	defer sched0.Stop()
	defer sched0.Wait()

	start := time.Now().UTC()
	expectedOrder := []string{}

	for i := 0; i < nTasks; i++ {
		key := fmt.Sprintf("task-%d", i)
		task := NewCustomTask(key, start.Add(time.Millisecond*10*time.Duration(i)))
		expectedOrder = append(expectedOrder, key)
		future := sched0.Submit(task)
		futures[i] = future
	}
	for i := range futures {
		fmt.Println(futures[i].Get())
	}
	if !reflect.DeepEqual(expectedOrder, O.Get()) {
		t.Errorf("execution order wrong, got: %v", O.Get())
	}
}

// O order
var O = Order{}

// Order is used for recording execution order
type Order struct {
	mu    sync.Mutex
	order []string
	first time.Time
	last  time.Time
}

// Push an execution id
func (o *Order) Push(s string) {
	o.mu.Lock()
	o.order = append(o.order, s)
	o.mu.Unlock()
}

// IsFirstZero check if first is zero time
func (o *Order) IsFirstZero() bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.first.IsZero()
}

// SetFirst time
func (o *Order) SetFirst(t time.Time) {
	o.mu.Lock()
	o.first = t
	o.mu.Unlock()
}

// GetFirst time
func (o *Order) GetFirst() time.Time {
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.first
}

// SetLast time
func (o *Order) SetLast(t time.Time) {
	o.mu.Lock()
	o.last = t
	o.mu.Unlock()
}

// GetLast time
func (o *Order) GetLast() time.Time {
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.last
}

// Get order
func (o *Order) Get() []string {
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.order
}

// Clear the order
func (o *Order) Clear() {
	o.mu.Lock()
	o.order = []string{}
	o.first = time.Time{}
	o.last = time.Time{}
	o.mu.Unlock()
}
