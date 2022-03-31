// Copyright 2022 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

// Map turns a []T1 to a []T2 using a mapping function.
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
func Reduce[T1, T2 any](s []T1, init T2, f func(T2, T1) T2) T2 {
	r := init
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter filters a slice by a given filter function f.
func Filter[T any](s []T, f func(T) bool) []T {
	r := make([]T, len(s))
	for _, v := range s {
		if !f(v) {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	x1 := []int{1, 2, 3, 4, 5}
	x2 := []float64{1, 2, 3, 4, 5}
	x3 := []byte{1, 2, 3, 4, 5}

	r1 := Reduce(Filter(Map(x1, func(t int) int {
		return 2 * t
	}), func(t int) bool {
		return t%3 == 0
	}), 0, func(t1, t2 int) int {
		return t1 + t2
	})
	r2 := Reduce(Filter(Map(x2, func(t float64) float64 {
		return 2 * t
	}), func(t float64) bool {
		return int(t)%3 == 0
	}), 0, func(t1, t2 float64) float64 {
		return t1 + t2
	})
	r3 := Reduce(Filter(Map(x3, func(t byte) byte {
		return 2 * t
	}), func(t byte) bool {
		return t%3 == 0
	}), 0, func(t1, t2 byte) byte {
		return t1 + t2
	})
	fmt.Println(r1, r2, r3)
}
