// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// List is a linked list.
type List[T any] struct {
	head, tail *element[T]
}

// An element is an entry in a linked list.
type element[T any] struct {
	next *element[T]
	val  T
}

// Push pushes an element to the end of the list.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v }
		lst.tail = lst.tail.next
	}
}

// Iterator ranges over a list.
type Iterator[T any] struct {
	next **element[T]
}

// Range returns an Iterator starting at the head of the list.
func (lst *List[T]) Range() *Iterator[T] {
	return &Iterator[T]{next: &lst.head}
}

// Next advances the iterator.
// It returns whether there are more elements.
func (it *Iterator[T]) Next() bool {
	if *it.next == nil {
		return false
	}
	it.next = &(*it.next).next
	return true
}

// Val returns the value of the current element.
// The bool result reports whether the value is valid.
func (it *Iterator[T]) Val() (T, bool) {
	if *it.next == nil {
		var zero T
		return zero, false
	}
	return (*it.next).val, true
}

// Transform runs a transform function on a list returning a new list.
func Transform[T1, T2 any](lst *List[T1], f func(T1) T2) *List[T2] {
	ret := &List[T2]{}
	it := lst.Range()
	for {
		if v, ok := it.Val(); ok {
			ret.Push(f(v))
		}
		it.Next()
	}
	return ret
}