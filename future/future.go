// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package future

import (
	"runtime"
	"sync/atomic"
)

type Future[T any] struct {
	value atomic.Value
}

// Get implements TaskFuture interface
func (f *Future[T]) Get() T {
	var v any
	for ; v == nil; v = f.value.Load() {
		runtime.Gosched()
	}
	return v.(T)
}

func (f *Future[T]) Put(v T) {
	f.value.Store(any(v))
}
