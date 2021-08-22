// Copyright 2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package slices defines various functions useful with slices of any type.
// Unless otherwise specified, these functions all apply to the elements
// of a slice at index 0 <= i < len(s).
package slices

import "golang.design/x/go2generics/constraints" // See #45458

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// EqualFunc reports whether two slices are equal using a comparison
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// index order, and the comparison stops at the first index for which
// eq returns false.
func EqualFunc[T1, T2 any](s1 []T1, s2 []T2, eq func(T1, T2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if !eq(s1[i], s2[i]) {
			return false
		}
	}
	return true
}

// Compare compares the elements of s1 and s2.
// The elements are compared sequentially starting at index 0,
// until one element is not equal to the other. The result of comparing
// the first non-matching elements is the result of the comparison.
// If both slices are equal until one of them ends, the shorter slice is
// considered less than the longer one
// The result will be 0 if s1==s2, -1 if s1 < s2, and +1 if s1 > s2.
func Compare[T constraints.Ordered](s1, s2 []T) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	l := min(len(s1), len(s2))

	for i := 0; i < l; i++ {
		switch {
		case s1[i] == s2[i]:
			continue
		case s1[i] < s2[i]:
			return -1
		case s1[i] > s2[i]:
			return 1
		}
	}
	switch {
	case l == len(s1) && l == len(s2) :
		return 0
	case l < len(s1): // s1 is longer
		return 1
	case l < len(s2): // s2 is longer
		fallthrough
	default:
		return -1
	}
}

// CompareFunc is like Compare, but uses a comparison function
// on each pair of elements. The elements are compared in index order,
// and the comparisons stop after the first time cmp returns non-zero.
// The result will be the first non-zero result of cmp; if cmp always
// returns 0 the result is 0 if len(s1) == len(s2), -1 if len(s1) < len(s2),
// and +1 if len(s1) > len(s2).
func CompareFunc[T any](s1, s2 []T, cmp func(T, T) int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	l := min(len(s1), len(s2))

	for i := 0; i < l; i++ {
		switch cmp(s1[i], s2[i]) {
		case 0:
			continue
		case -1:
			return -1
		case 1:
			return 1
		}
	}
	switch {
	case l == len(s1) && l == len(s2) :
		return 0
	case l < len(s1): // s1 is longer
		return 1
	case l < len(s2): // s2 is longer
		fallthrough
	default:
		return -1
	}
}

// Index returns the index of the first occurrence of v in s, or -1 if not present.
func Index[T comparable](s []T, v T) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// IndexFunc returns the index into s of the first element
// satisfying f(c), or -1 if none do.
func IndexFunc[T any](s []T, f func(T) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

// Contains reports whether v is present in s.
func Contains[T comparable](s []T, v T) bool {
	return Index[T](s, v) != -1
}

// Insert inserts the values v... into s at index i, returning the modified slice.
// In the returned slice r, r[i] == the first v.  Insert panics if i is out of range.
// This function is O(len(s) + len(v)).
func Insert[S constraints.Slice[T], T any](s S, i int, v ...T) S {
	if i > len(s) {
		panic("slices: out of slice index")
	}
	s = s[:i]
	s = append(s, v...)
	return s
}

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if s[i:j] is not a valid slice of s.
// Delete modifies the contents of the slice s; it does not create a new slice.
// Delete is O(len(s)-(j-i)), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
func Delete[S constraints.Slice[T], T any](s S, i, j int) S {
	if !(0 < i && i < len(s)) || !(0 < j && j < len(s)) || i > j {
		panic("slices: invalid index i or j")
	}

	s = s[:i]
	s = append(s, s[j:]...)
	return s
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func Clone[S constraints.Slice[T], T any](s S) S {
	ss := make([]T, len(s))
	for i := range s {
		ss[i] = s[i]
	}
	return ss
}

// Compact replaces consecutive runs of equal elements with a single copy.
// This is like the uniq command found on Unix.
// Compact modifies the contents of the slice s; it does not create a new slice.
func Compact[S constraints.Slice[T], T comparable](s S) S {
	if s == nil {
		return s
	}
	lastIdx := len(s)-1
	for i := len(s)-2; i !=0; i-- {
		if s[i] == s[lastIdx] {
			continue
		}
		s = append(s[:i+1], s[lastIdx:]...)
		lastIdx = i
	}
	return s
}

// CompactFunc is like Compact, but uses a comparison function.
func CompactFunc[S constraints.Slice[T], T any](s S, cmp func(T, T) bool) S {
	if s == nil {
		return s
	}
	lastIdx := len(s)-1
	for i := len(s)-2; i !=0; i-- {
		if cmp(s[i], s[lastIdx]) {
			continue
		}
		s = append(s[:i+1], s[lastIdx:]...)
		lastIdx = i
	}
	return s
}

// Grow grows the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow will panic.
func Grow[S constraints.Slice[T], T any](s S, n int) S {
	ss := make([]T, len(s), len(s) + n)
	copy(ss, s[:])
	return ss
}

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
func Clip[S constraints.Slice[T], T any](s S) S {
	s = s[:len(s):len(s)]
	return s
}