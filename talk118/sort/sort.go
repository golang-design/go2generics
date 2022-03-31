// Copyright 2022 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"reflect"
	"sort"
)

type wrapSort[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func (s wrapSort[T]) Len() int           { return len(s.s) }
func (s wrapSort[T]) Less(i, j int) bool { return s.cmp(s.s[i], s.s[j]) }
func (s wrapSort[T]) Swap(i, j int)      { s.s[i], s.s[j] = s.s[j], s.s[i] }

// Sort wraps sort.Sort
func Sort[T any](s []T, cmp func(T, T) bool) {
	sort.Sort(wrapSort[T]{s, cmp})
}

// QuickSort implements a generic quick sort.
func QuickSort[T comparable](s []T, cmp func(a, b T) bool) {
	quickSort(s, 0, len(s)-1, cmp)
}
func quickSort[T comparable](s []T, start, end int, cmp func(a, b T) bool) {
	if end-start < 1 {
		return
	}
	pivot := s[end]
	split := start
	for i := start; i < end; i++ {
		if cmp(s[i], pivot) {
			s[i], s[split] = s[split], s[i]
			split++
		}
	}
	s[end] = s[split]
	s[split] = pivot
	quickSort(s, start, split-1, cmp)
	quickSort(s, split+1, end, cmp)
}

type Tests[T any] struct {
	Orig []T
	Want []T
}

func main() {
	testsInt := []Tests[int]{
		{
			Orig: []int{2, 3, 1},
			Want: []int{1, 2, 3},
		},
	}

	for _, t := range testsInt {
		Sort(t.Orig, func(a, b int) bool {
			return a < b
		})
		if !reflect.DeepEqual(t.Orig, t.Want) {
			panic(fmt.Sprintf("want %v, got %v", t.Want, t.Orig))
		}
	}

	testsFloat := []Tests[float32]{
		{
			Orig: []float32{2, 3, 1},
			Want: []float32{1, 2, 3},
		},
	}

	for _, t := range testsFloat {
		Sort(t.Orig, func(a, b float32) bool {
			return a < b
		})
		if !reflect.DeepEqual(t.Orig, t.Want) {
			panic(fmt.Sprintf("want %v, got %v", t.Want, t.Orig))
		}
	}

	println("OK")
}
