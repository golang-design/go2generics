// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slices

import (
	"sort"

	"golang.design/x/go2generics/math"
)

// orderedSlice is a slice of values of some ordered type.
type orderedSlice[Elem math.Ordered] []Elem

// orderedSlice implements sort.Interface.

func (s orderedSlice[Elem]) Len() int           { return len(s) }
func (s orderedSlice[Elem]) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	}
	isNaN := func(f Elem) bool { return f != f }
	if isNaN(s[i]) && !isNaN(s[j]) {
		return true
	}
	return false
}
func (s orderedSlice[Elem]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// SortOrderedSlice sorts a slice of any ordered type in ascending order.
func SortOrderedSlice[Elem math.Ordered](s []Elem) {
	sort.Sort(orderedSlice[Elem](s))
}

// sliceFn implements sort.Interface for a slice of any type with an
// explicit less-than function.
type sliceFn[Elem any] struct {
	s    []Elem
	less func(Elem, Elem) bool
}

func (s sliceFn[Elem]) Len() int           { return len(s.s) }
func (s sliceFn[Elem]) Less(i, j int) bool { return s.less(s.s[i], s.s[j]) }
func (s sliceFn[Elem]) Swap(i, j int)      { s.s[i], s.s[j] = s.s[j], s.s[i] }

// SliceFn sorts a slice of any type according to a less-than function.
func SliceFn[Elem any](s []Elem, less func(Elem, Elem) bool) {
	sort.Sort(sliceFn[Elem]{s, less})
}
