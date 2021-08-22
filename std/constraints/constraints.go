// Copyright 2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package constraints defines a set of useful constraints to be used with type parameters.
package constraints

// Signed is a constraint that permits any signed integer type.
type Signed interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

// Unsigned is a constraint that permits any unsigned integer type.
type Unsigned interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr
}

// Integer is a constraint that permits any integer type.
type Integer interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int |
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint |
	~uintptr
}

// Float is a constraint that permits any floating-point type.
type Float interface {
	~float32 | ~float64
}

// Complex is a constraint that permits any complex numeric type.
type Complex interface {
	~complex64 | ~complex128
}

// Ordered is a constraint that permits any ordered type: any type that supports the operators < <= >= >.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | 
	~uintptr | ~float32 | ~float64 | ~string
}

// Slice is a constraint that permits any slice.
type Slice[T any] interface { ~[]T }

// Map is a constraint that permits any map.
type Map[K comparable, V any] interface { ~map[K]V }

// Chan is a constraint that permits any channel.
type Chan[T any] interface { ~chan T }