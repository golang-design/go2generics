// Copyright 2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package maps defines various functions useful with maps of any type.
package maps

// Keys returns the keys of the map m.
// The keys will be an indeterminate order.
func Keys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func Values[K comparable, V any](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// Equal reports whether two maps contain the same key/value pairs.
// Values are compared using ==.
func Equal[K, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// EqualFunc is like Equal, but compares values using cmp.
// Keys are still compared with ==.
func EqualFunc[K comparable, V1, V2 any](m1 map[K]V1, m2 map[K]V2, cmp func(V1, V2) bool) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || !cmp(v1, v2) {
			return false
		}
	}
	return true
}

// Clear removes all entries from m, leaving it empty.
func Clear[K comparable, V any](m map[K]V) {
	for k := range m {
		delete(m, k)
	}
}

// Clone returns a copy of m.  This is a shallow clone:
// the new keys and values are set using ordinary assignment.
func Clone[K comparable, V any](m map[K]V) map[K]V {
	mm := map[K]V{}
	for k, v := range m {
		mm[k] = v
	}
	return mm
}

// Add adds all key/value pairs in src to dst. When a key in src
// is already present in dst, the value in dst will be overwritten
// by the value associated with the key in src.
func Add[K comparable, V any](dst, src map[K]V) {
	for k, v := range src {
		dst[k] = v
	}
}

// Filter deletes any key/value pairs from m for which keep returns false.
func Filter[K comparable, V any](m map[K]V, keep func(K, V) bool) {
	for k, v := range m {
		if !keep(k, v) {
			delete(m, k)
		}
	}
}