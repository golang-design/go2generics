// Code generated by go2go; DO NOT EDIT.


//line sched.go2:1
package sched

//line sched.go2:1
import (
//line sched.go2:1
 "container/heap"
//line sched.go2:1
 "context"
//line sched.go2:1
 "fmt"
//line sched.go2:1
 "reflect"
//line sched.go2:1
 "runtime"
//line sched.go2:1
 "sync"
//line sched.go2:1
 "sync/atomic"
//line sched.go2:1
 "testing"
//line sched.go2:1
 "time"
//line sched.go2:1
 "unsafe"
//line sched.go2:1
)

//line sched.go2:1
type Importable୦ int

//line sched.go2:1
var _ = heap.Fix
//line sched.go2:1
var _ = context.Background
//line sched.go2:1
var _ = fmt.Errorf
//line sched.go2:1
var _ = reflect.Append
//line sched.go2:1
var _ = runtime.BlockProfile

//line sched.go2:1
type _ sync.Cond

//line sched.go2:1
var _ = atomic.AddInt32
//line sched.go2:1
var _ = testing.AllocsPerRun

//line sched.go2:1
const _ = time.ANSIC

//line sched.go2:1
type _ unsafe.Pointer
