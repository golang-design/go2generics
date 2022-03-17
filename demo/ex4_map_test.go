// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package demo_test

import "testing"

// naive generic map[k]v
type Pair[T1, T2 any] struct {
	Key   T1
	Value T2
}
type GenericMap[T1 comparable, T2 any] struct {
	s []Pair[T1, T2]
}

func NewGenericMap[T1 comparable, T2 any]() GenericMap[T1, T2] {
	return GenericMap[T1, T2]{s: [](Pair[T1, T2]){}}
}
func (m *GenericMap[T1, T2]) Set(k T1, v T2) {
	m.s = append(m.s, Pair[T1, T2]{k, v})
}
func (m *GenericMap[T1, T2]) Get(k T1) (v T2, ok bool) {
	for _, p := range m.s {
		if p.Key == k {
			return p.Value, true
		}
	}
	return
}

func TestGenericMap(t *testing.T) {
	m := NewGenericMap[int, float64]()
	m.Set(1, 2.0)
	m.Set(2, 3.0)
	m.Set(3, 4.0)

	if v, ok := m.Get(1); !ok || v != 2.0 {
		panic("get 1 fail")
	}
	if v, ok := m.Get(2); !ok || v != 3.0 {
		panic("get 2 fail")
	}
	if v, ok := m.Get(3); !ok || v != 4.0 {
		panic("get 3 fail")
	}
	println("OK")
}
