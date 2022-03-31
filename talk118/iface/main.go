// Copyright 2022 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func foo[T any](x T) T { return x }

type Ia[T any] interface{ *T }
type Ib[T any] interface{ Foo() }

// func bar(T Ia[int]) {} // ERROR: interface contains type constraints
// func bar(T Ib[int]) {} // OK

func Foo[T any]() {}

func main() {
	foo("string")
	// _ = Foo // ERROR: cannot use generic function Foo without instantiation
	_ = Foo[int]
}

func foo1[T fmt.Stringer](t T) string {
	return t.String()
}
func foo2(t fmt.Stringer) string {
	return t.String()
}
