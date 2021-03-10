// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slices

import "golang.design/x/go2generics/math"

// Equal reports whether two slices are equal
func Equal[Elem math.Ordered](s1, s2 []Elem) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if v1 != v2 {
			if !math.IsNaN(v1) || !math.IsNaN(v2) {
				return false
			}
		}
	}
	return true
}

// EqualFn reports whether two slices are equal using a comparision
// function on each element.
func EqualFn[Elem any](s1, s2 []Elem, eq func(Elem, Elem) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if !eq(v1, v2) {
			return false
		}
	}
	return true
}

// Map turns a []T1 to a []T2 using a mapping function.
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter filters values from a slice using a filter function.
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Reverse is a function that takes a []T argument and
// reverses that slice in place.
func Reverse[T any](list []T) {
	i := 0
	j := len(list)-1
	for i < j {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
}
