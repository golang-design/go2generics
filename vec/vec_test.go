// Copyright 2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vec

import (
	"testing"
)

func TestVec_New(t *testing.T) {
	t.Run("Vec2", func(t *testing.T) {
		v1 := NewVec2[float64](1, 1)
		v2 := NewVec2[float64](2, 2)
		v3 := NewVec2[float64](1, 1)
		if v1.Eq(v2) {
			t.Fatalf("unexpected comparison, got true, want false")
		}
		if !v1.Eq(v3) {
			t.Fatalf("unexpected comparison, got false, want true")
		}
	})
}

func TestVec_Rand(t *testing.T) {
	t.Run("Vec2", func(t *testing.T) {
		v1 := NewRandVec2[float64]()
		v2 := NewRandVec2[float64]()
		if v1.Eq(v2) {
			t.Fatalf("unexpected different random vectors, got %v, %v", v1, v2)
		}
	})
}

func TestVec_Add(t *testing.T) {
	t.Run("Vec2", func(t *testing.T) {
		v1 := NewVec2[float64](1, 1)
		v2 := NewVec2[float64](2, 2)
		got := v1.Add(v2)
		want := NewVec2[float64](3, 3)
		if !want.Eq(got) {
			t.Fatalf("unexpected Add, got %v, want %v", got, want)
		}
	})
}

func TestVec_Sub(t *testing.T) {
	t.Run("Vec2", func(t *testing.T) {
		v1 := NewVec2[float64](1, 1)
		v2 := NewVec2[float64](2, 2)
		got := v1.Sub(v2)
		want := NewVec2[float64](-1, -1)
		if !want.Eq(got) {
			t.Fatalf("unexpected Add, got %v, want %v", got, want)
		}
	})
}

func TestVec_Dot(t *testing.T) {
	t.Run("Vec2", func(t *testing.T) {
		v1 := NewVec2[float64](1, 1)
		v2 := NewVec2[float64](2, 2)
		got := v1.Dot(v2)
		want := 4.0
		if !ApproxEq(want, got, Epsilon) {
			t.Fatalf("unexpected Add, got %v, want %v", got, want)
		}
	})
}

func TestVec_IsZero(t *testing.T) {
	t.Run("Vec2", func(t *testing.T) {
		v1 := NewVec2[float64](1, 1)
		if v1.IsZero() {
			t.Fatalf("unexpected IsZero assertion, want false, got true")
		}
		v1 = NewVec2[float64](0, 0)
		if !v1.IsZero() {
			t.Fatalf("unexpected IsZero assertion, want true, got false")
		}
	})
}

func TestVec_Scale(t *testing.T) {
	t.Run("Vec2", func(t *testing.T) {
		v1 := NewVec2[float64](1, 1)
		got := v1.Scale(2, 2)
		want := NewVec2[float64](2, 2)
		if !got.Eq(want) {
			t.Fatalf("unexpected scale, want %v, got %v", want, got)
		}
	})
}

func TestVec_Translate(t *testing.T) {
	t.Run("Vec2", func(t *testing.T) {
		v1 := NewVec2[float64](1, 1)
		got := v1.Translate(2, 2)
		want := NewVec2[float64](3, 3)
		if !got.Eq(want) {
			t.Fatalf("unexpected translate, want %v, got %v", want, got)
		}
	})
}
