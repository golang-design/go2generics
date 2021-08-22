// Copyright 2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linalg

import "testing"

func TestDotProduct(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	val := DotProduct(a, b)
	if val != 20 {
		t.Fatalf("unexpected DotProduct result, got %v want %v", val, 20)
	}
}