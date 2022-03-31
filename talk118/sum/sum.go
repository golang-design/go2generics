// Copyright 2022 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type T interface {
	Add(T) T
}

type MyInt int

func (i MyInt) Add(val T) T {
	switch v := val.(type) {
	case MyInt:
		return T(i + v)
	default:
		panic("unsupported T")
	}
}

func Sum(elems ...T) (sum T) { // T could be any type that implements Add()
	if len(elems) == 0 {
		return
	}
	sum = elems[0]
	for _, v := range elems[1:] {
		sum = sum.Add(v)
	}
	return
}

func f1() {
	vals := []MyInt{1, 2, 3, 4}
	sum := MyInt(0)
	for _, v := range vals {
		sum = Sum(sum, v).(MyInt)
	}
	println(sum)
}

func GenericSum[T ~int](elems ...T) (sum T) { // The underlying type of T must be int
	for i := range elems {
		sum += elems[i]
	}
	return
}

func f2() {
	vals := []MyInt{1, 2, 3, 4}
	sum := GenericSum(vals...)
	println(sum)
}

func main() {
	f1()
	f2()
}
