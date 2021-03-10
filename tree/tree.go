// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tree

// A generic binary tree might be declared as follows.
type Tree[E any] struct {
	left, right *Tree[E]
	payload E
}