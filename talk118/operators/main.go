// Copyright 2022 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | // Signed
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | // Unsigned
		~float32 | ~float64 | // Float
		~string
}

// IsSorted reports whether x is sorted in ascending order.
func IsSorted[E Ordered](x []E) bool {
	for i := len(x) - 1; i > 0; i-- {
		if x[i] < x[i-1] {
			return false
		}
	}
	return true
}

// Incorrect
// type Comparable[T any] interface {
// 	==(T) bool
// }

type Comparable[T any] interface {
	Equal(T) bool
}

var (
	// cannot use MyInt(0) (constant 0 of type MyInt) as Comparable[int]
	// value in variable declaration: MyInt does not implement Comparable[int]
	// (wrong type for method Equal)
	// 	have Equal(j MyInt) bool
	// 	want Equal(int) bool
	// _ Comparable[int] = MyInt(0)
	_ Comparable[int] = MyInt2(0)
)

type MyInt int

func (i MyInt) Equal(j MyInt) bool {
	return i == j
}

type MyInt2 int

func (i MyInt2) Equal(j int) bool {
	return int(i) == j
}
