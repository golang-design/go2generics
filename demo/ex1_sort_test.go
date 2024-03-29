// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package demo_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type wrapSort[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func (s wrapSort[T]) Len() int           { return len(s.s) }
func (s wrapSort[T]) Less(i, j int) bool { return s.cmp(s.s[i], s.s[j]) }
func (s wrapSort[T]) Swap(i, j int)      { s.s[i], s.s[j] = s.s[j], s.s[i] }

func Sort[T any](s []T, cmp func(T, T) bool) {
	sort.Sort(wrapSort[T]{s, cmp})
}

func TestSort(t *testing.T) {
	s1 := []int{2, 3, 1}
	sorted1 := []int{1, 2, 3}

	Sort(s1, func(a, b int) bool {
		return a < b
	})

	if !reflect.DeepEqual(s1, sorted1) {
		panic(fmt.Sprintf("want %v, got %v", sorted1, s1))
	}
	fmt.Printf("sorted: %v\n", s1)

	s2 := []float64{2.0, 3.0, 1.0}
	sorted2 := []float64{1, 2, 3}

	Sort(s2, func(a, b float64) bool {
		return a < b
	})

	if !reflect.DeepEqual(s2, sorted2) {
		panic(fmt.Sprintf("want %v, got %v", sorted2, s2))
	}
	fmt.Printf("sorted: %v\n", s2)
}
