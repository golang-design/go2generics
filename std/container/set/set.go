// Copyright 2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package set defines a Set type that holds a set of elements.
package set

// A Set is a set of elements of some comparable type.
// Sets are implemented using maps, and have similar performance characteristics.
// Like maps, Sets are reference types.
// That is, for Sets s1 = s2 will leave s1 and s2 pointing to the same set of elements:
// changes to s1 will be reflected in s2 and vice-versa.
// Unlike maps, the zero value of a Set is usable; there is no equivalent to make.
// As with maps, concurrent calls to functions and methods that read values are fine;
// concurrent calls to functions and methods that write values are racy.
type Set[Elem comparable] struct {
	m map[Elem]struct{}
}


// Of returns a new set containing the listed elements.
func Of[Elem comparable](v ...Elem) Set[Elem] {
	m := map[Elem]struct{}{}
	for _, e := range v {
		m[e] = struct{}{}
	}
	return Set[Elem]{m}
}

// Add adds elements to a set.
func (s *Set[Elem]) Add(v ...Elem) {
	for _, e := range v {
		s.m[e] = struct{}{}
	}
}

// AddSet adds the elements of set s2 to s.
func (s *Set[Elem]) AddSet(s2 Set[Elem]) {
	for k := range s2.m {
		s.m[k] = struct{}{}
	}
}

// Remove removes elements from a set.
// Elements that are not present are ignored.
func (s *Set[Elem]) Remove(v ...Elem) {
	for k := range s.m {
		for i := range v {
			if k == v[i] {
				delete(s.m, v[i])
			}
		}
	}
}

// RemoveSet removes the elements of set s2 from s.
// Elements present in s2 but not s are ignored.
func (s *Set[Elem]) RemoveSet(s2 Set[Elem]) {
	for k1 := range s.m {
		for k2 := range s2.m {
			if k1 == k2 {
				delete(s.m, k1)
			}
		}
	}
}

// Has reports whether v is in the set.
func (s *Set[Elem]) Has(v Elem) bool {
	_, ok := s.m[v]
	return ok
}

// HasAny reports whether any of the elements in s2 are in s.
func (s *Set[Elem]) HasAny(s2 Set[Elem]) bool {
	for k2 := range s2.m {
		if _, ok := s.m[k2]; ok {
			return true
		}
	}
	return false
}

// HasAll reports whether all of the elements in s2 are in s.
func (s *Set[Elem]) HasAll(s2 Set[Elem]) bool {
	for k2 := range s2.m {
		if _, ok := s.m[k2]; !ok {
			return false
		}
	}
	return true
}

// Values returns the elements in the set s as a slice.
// The values will be in an indeterminate order.
func (s *Set[Elem]) Values() []Elem {
	sli := make([]Elem, len(s.m))
	for k := range s.m {
		sli = append(sli, k)
	}
	return sli
}

// Equal reports whether s and s2 contain the same elements.
func (s *Set[Elem]) Equal(s2 Set[Elem]) bool {
	if len(s.m) != len(s2.m) {
		return false
	}

	for k := range s.m {
		if _, ok := s2.m[k]; !ok {
			return false
		}
	}
	return true
}

// Clear removes all elements from s, leaving it empty.
func (s *Set[Elem]) Clear() {
	s.m = map[Elem]struct{}{}
}

// Clone returns a copy of s.
// The elements are copied using assignment,
// so this is a shallow clone.
func (s *Set[Elem]) Clone() Set[Elem] {
	m := map[Elem]struct{}{}
	for k := range s.m {
		m[k] = struct{}{}
	}
	return Set[Elem]{m}
}

// Filter deletes any elements from s for which keep returns false.
func (s *Set[Elem]) Filter(keep func(Elem) bool) {
	for k := range s.m {
		if !keep(k) {
			delete(s.m, k)
		}
	}
}

// Len returns the number of elements in s.
func (s *Set[Elem]) Len() int {
	return len(s.m)
}

// Do calls f on every element in the set s,
// stopping if f returns false.
// f should not change s.
// f will be called on values in an indeterminate order.
func (s *Set[Elem]) Do(f func(Elem) bool) {
	for k := range s.m {
		if !f(k) {
			return
		}
	}
}

// Union constructs a new set containing the union of s1 and s2.
func Union[Elem comparable](s1, s2 Set[Elem]) Set[Elem] {
	m := map[Elem]struct{}{}
	for k := range s1.m {
		m[k] = struct{}{}
	}
	for k := range s2.m {
		m[k] = struct{}{}
	}
	return Set[Elem]{m}
}

// Intersection constructs a new set containing the intersection of s1 and s2.
func Intersection[Elem comparable](s1, s2 Set[Elem]) Set[Elem] {
	m := map[Elem]struct{}{}
	for k := range s1.m {
		if _, ok := s2.m[k]; ok {
			m[k] = struct{}{}
		}
	}
	return Set[Elem]{m}
}

// Difference constructs a new set containing the elements of s1 that
// are not present in s2.
func Difference[Elem comparable](s1, s2 Set[Elem]) Set[Elem] {
	m := map[Elem]struct{}{}
	for k := range s1.m {
		if _, ok := s2.m[k]; !ok {
			m[k] = struct{}{}
		}
	}
	return Set[Elem]{m}
}