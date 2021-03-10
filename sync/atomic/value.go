// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import "sync/atomic"

// Value is a type-safe atomic value
type Value[T any] struct {
	val atomic.Value
}

func (v *Value[T]) Load() T {
	x := v.val.Load()
	return x.(T)
}

func (v *Value[T]) Store(val T) {
	v.val.Store(interface{}(val))
}
