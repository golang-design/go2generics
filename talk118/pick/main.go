// Copyright 2022 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math/rand"
)

func Pick[S ~[]Elem, Elem any](s S) Elem {
	return s[rand.Intn(len(s))]
}

func Pick2[S ~[]any](s S) any {
	return s[rand.Intn(len(s))]
}

func toAny[T any](s []T) []any {
	ret := make([]any, len(s))
	for i, v := range s {
		ret[i] = v
	}
	return ret
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	x := Pick(s)
	// y := Pick2(s) // ERROR: []int does not implement ~[]any
	y := Pick2(toAny(s))
	// x = y // ERROR: cannot use y (variable of type any) as int value in assignment: need type assertion

	fmt.Printf("%T, %T", x, y)
}
