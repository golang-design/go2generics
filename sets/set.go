// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sets

type Set[Elem comparable] map[Elem]struct{}

func Make[Elem comparable]() Set[Elem] {
	return make(Set[Elem])
}

func (s Set[Elem]) Add(v Elem) {
	s[v] = struct{}{}
}

func (s Set[Elem]) Delete(v Elem) {
	delete(s, v)
}

func (s Set[Elem]) Contains(v Elem) bool {
	_, ok := s[v]
	return ok
}

func (s Set[Elem]) Len() int {
	return len(s)
}

func (s Set[Elem]) Iterate(f func(Elem)) {
	for v := range s { f(v) }
}

func (s Set[Elem]) Values() []Elem {
	r := make([]Elem, 0, len(s))
	for v := range s {
		r = append(r, v)
	}
	return r
}

func Equal[Elem comparable](s1, s2 Set[Elem]) bool {
	if len(s1) != len(s2) {
		return false
	}
	for v1 := range s1 {
		if !s2.Contains(v1) {
			return false
		}
	}
	return true
}

func (s Set[Elem]) Copy() Set[Elem] {
	r := Set[Elem](make(map[Elem]struct{}, len(s)))
	for v := range s {
		r[v] = struct{}{}
	}
	return r
}

func (s Set[Elem]) AddSet(s2 Set[Elem]) {
	for v := range s2 {
		s[v] = struct{}{}
	}
}

func (s Set[Elem]) SubSet(s2 Set[Elem]) {
	for v := range s2 {
		delete(s, v)
	}
}

func (s Set[Elem]) Intersect(s2 Set[Elem]) {
	for v := range s {
		if !s2.Contains(v) {
			delete(s, v)
		}
	}
}

func (s Set[Elem]) Filter(f func(Elem) bool) {
	for v := range s {
		if !f(v) {
			delete(s, v)
		}
	}
}