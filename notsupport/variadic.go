// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notsupport

// This file illustrates what's impossible in the current
// contract design - variadic generics.

// type Tuple [Ts ...comparable] struct {
// 	elements ...Ts
// }

// func (t *Tuple(Ts...)) Set(es ...Ts) {
// 	t.elements(Ts...){es...}
// }

// func (t Tuple) PirntAll() {
// 	for _, e := range t.elements {
// 		fmt.Println(e)
// 	}
// }

// // func (t Tuple(Ts...)) Get(i int) T ??
