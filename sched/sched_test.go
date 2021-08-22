package sched

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

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
