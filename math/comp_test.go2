// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import "testing"

func TestMax(t *testing.T) {
	got := Max(1, 2, 3)
	if got != 3 {
		t.Fatalf("want 3")
	}
	got = Max(1, 3, 2)
	if got != 3 {
		t.Fatalf("want 3")
	}
	got = Max(2, 1, 3)
	if got != 3 {
		t.Fatalf("want 3")
	}
	got = Max(2, 3, 1)
	if got != 3 {
		t.Fatalf("want 3")
	}
	got = Max(3, 1, 2)
	if got != 3 {
		t.Fatalf("want 3")
	}
	got = Max(3, 2, 1)
	if got != 3 {
		t.Fatalf("want 3")
	}
}

func TestMin(t *testing.T) {
	m := Min(1, 2, 3)
	if m != 1 {
		t.Fatalf("want 3, got %v", m)
	}
}
