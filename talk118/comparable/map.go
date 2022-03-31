// Copyright 2022 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[E comparable](s []E, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

func Equal[T comparable](v1, v2 T) bool {
	return v1 == v2
}

func main() {
	// v1 := interface{}(func() {})
	// v2 := interface{}(func() {})
	// Equal(v1, v2)
}

type P map[interface{}]struct{}     // OK
type R[T comparable] map[T]struct{} // OK
// type Q[T any] map[T]struct{}     // ERROR: incomparable map key type T
