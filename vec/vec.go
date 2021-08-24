// Copyright 2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vec

import (
	"math"
	"math/rand"
)

const (
	// Epsilon is a default epsilon value for computation.
	Epsilon = 1e-7
)

// Vec2 represents a 2D vector (x, y).
type Vec2 struct {
	X, Y float64
}

// NewVec2 creates a 2D vector with given parameters.
func NewVec2(x, y float64) Vec2 {
	return Vec2{x, y}
}

// NewRandVec2 returns a random 2D vector where all components are
// sitting in range [0, 1].
func NewRandVec2() Vec2 {
	return Vec2{
		float64(rand.Float64()),
		float64(rand.Float64()),
	}
}

// Eq compares the two given vectors, and returns true if they are equal.
func (v Vec2) Eq(u Vec2) bool {
	if ApproxEq(float64(v.X), float64(u.X), Epsilon) &&
		ApproxEq(float64(v.Y), float64(u.Y), Epsilon) {
		return true
	}
	return false
}

// Add add the two given vectors, and returns the resulting vector.
func (v Vec2) Add(u Vec2) Vec2 {
	return Vec2{v.X + u.X, v.Y + u.Y}
}

// Sub subtracts the two given vectors, and returns the resulting vector.
func (v Vec2) Sub(u Vec2) Vec2 {
	return Vec2{v.X - u.X, v.Y - u.Y}
}

// IsZero checks if the given vector is a zero vector.
func (v Vec2) IsZero() bool {
	if ApproxEq(float64(v.X), 0, Epsilon) &&
		ApproxEq(float64(v.Y), 0, Epsilon) {
		return true
	}
	return false
}

// Scale scales the given 2D vector and returns the resulting vector.
func (v Vec2) Scale(x, y float64) Vec2 {
	return Vec2{v.X * x, v.Y * y}
}

// float64ranslate translates the given 2D vector and returns the resulting vector.
func (v Vec2) Translate(x, y float64) Vec2 {
	return Vec2{v.X + x, v.Y + y}
}

// Dot computes the dot product of two given vectors.
func (v Vec2) Dot(u Vec2) float64 {
	return v.X*u.X + v.Y*u.Y
}

// Len returns the length of the given vector.
func (v Vec2) Len() float64 {
	return float64(math.Sqrt(float64(v.Dot(v))))
}

// Unit computes the unit vector along the direction of the given vector.
func (v Vec2) Unit() Vec2 {
	n := 1.0 / v.Len()
	return Vec2{v.X * n, v.Y * n}
}
