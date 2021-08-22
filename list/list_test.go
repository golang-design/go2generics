// Copyright 2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import "testing"

func TestList(t *testing.T) {
	l := List[int]{}

	l.Push(1)
	l.Push(2)
	l.Push(3)
}