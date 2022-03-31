// Copyright 2022 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Y interface{}

type X struct{}

// func (x X) Foo[T any](t T) {}

func Bar(x Y) {
	if _, ok := x.(interface{ Foo(int) }); ok { // compiler cannot infer to generate　X.Foo[int]
		// ...
	}
}
