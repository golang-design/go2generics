// Copyright 2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vec

import (
	"math/rand"
	"testing"
)

var (
	v  float64
	vv Vec2[float64]
)

func BenchmarkVec_Eq(b *testing.B) {
	b.Run("Vec2", func(b *testing.B) {
		v1 := NewVec2[float64](rand.Float64(), rand.Float64())
		v2 := NewVec2[float64](rand.Float64(), rand.Float64())

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = v1.Eq(v2)
		}
	})
}

func BenchmarkVec_Add(b *testing.B) {
	b.Run("Vec2", func(b *testing.B) {
		v1 := NewVec2[float64](rand.Float64(), rand.Float64())
		v2 := NewVec2[float64](rand.Float64(), rand.Float64())

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			vv = v1.Add(v2)
		}
	})
}

func BenchmarkVec_Sub(b *testing.B) {
	b.Run("Vec2", func(b *testing.B) {
		v1 := NewVec2[float64](rand.Float64(), rand.Float64())
		v2 := NewVec2[float64](rand.Float64(), rand.Float64())

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			vv = v1.Sub(v2)
		}
	})
}
func BenchmarkVec_IsZero(b *testing.B) {
	b.Run("Vec2", func(b *testing.B) {
		v1 := NewVec2[float64](rand.Float64(), rand.Float64())

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = v1.IsZero()
		}
	})
}

func BenchmarkVec_Scale(b *testing.B) {
	b.Run("Vec2", func(b *testing.B) {
		v1 := NewVec2[float64](rand.Float64(), rand.Float64())

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			vv = v1.Scale(2, 2)
		}
	})
}

func BenchmarkVec_Translate(b *testing.B) {
	b.Run("Vec2", func(b *testing.B) {
		v1 := NewVec2[float64](rand.Float64(), rand.Float64())

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			vv = v1.Translate(2, 2)
		}
	})
}

func BenchmarkVec_Dot(b *testing.B) {
	b.Run("Vec2", func(b *testing.B) {
		v1 := NewVec2[float64](rand.Float64(), rand.Float64())
		v2 := NewVec2[float64](rand.Float64(), rand.Float64())

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v = v1.Dot(v2)
		}
	})
}

func BenchmarkVec_Len(b *testing.B) {
	b.Run("Vec2", func(b *testing.B) {
		v1 := NewVec2[float64](rand.Float64(), rand.Float64())

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			v = v1.Len()
		}
	})
}

func BenchmarkVec_Unit(b *testing.B) {
	b.Run("Vec2", func(b *testing.B) {
		v1 := NewVec2[float64](rand.Float64(), rand.Float64())

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			vv = v1.Unit()
		}
	})
}
func BenchmarkVec_Apply(b *testing.B) {
	b.Run("Vec2", func(b *testing.B) {
		v1 := NewVec2[float64](rand.Float64(), rand.Float64())

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			vv = v1.Unit()
		}
	})
}
