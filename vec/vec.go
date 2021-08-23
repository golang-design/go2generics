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

type floats interface {
	float32 | float64
}

// Vec2 represents a 2D vector (x, y).
type Vec2[T floats] struct {
	X, Y T
}

// NewVec2 creates a 2D vector with given parameters.
func NewVec2[T floats](x, y T) Vec2[T] {
	return Vec2[T]{x, y}
}

// NewRandVec2 returns a random 2D vector where all components are
// sitting in range [0, 1].
func NewRandVec2[T floats]() Vec2[T] {
	return Vec2[T]{
		T(rand.Float64()),
		T(rand.Float64()),
	}
}

// Eq compares the two given vectors, and returns true if they are equal.
func (v Vec2[T]) Eq(u Vec2[T]) bool {
	if ApproxEq(float64(v.X), float64(u.X), Epsilon) &&
		ApproxEq(float64(v.Y), float64(u.Y), Epsilon) {
		return true
	}
	return false
}

// Add add the two given vectors, and returns the resulting vector.
func (v Vec2[T]) Add(u Vec2[T]) Vec2[T] {
	return Vec2[T]{v.X + u.X, v.Y + u.Y}
}

// Sub subtracts the two given vectors, and returns the resulting vector.
func (v Vec2[T]) Sub(u Vec2[T]) Vec2[T] {
	return Vec2[T]{v.X - u.X, v.Y - u.Y}
}

// IsZero checks if the given vector is a zero vector.
func (v Vec2[T]) IsZero() bool {
	if ApproxEq(float64(v.X), 0, Epsilon) &&
		ApproxEq(float64(v.Y), 0, Epsilon) {
		return true
	}
	return false
}

// Scale scales the given 2D vector and returns the resulting vector.
func (v Vec2[T]) Scale(x, y T) Vec2[T] {
	return Vec2[T]{v.X * x, v.Y * y}
}

// Translate translates the given 2D vector and returns the resulting vector.
func (v Vec2[T]) Translate(x, y T) Vec2[T] {
	return Vec2[T]{v.X + x, v.Y + y}
}

// Dot computes the dot product of two given vectors.
func (v Vec2[T]) Dot(u Vec2[T]) T {
	return v.X*u.X + v.Y*u.Y
}

// Len returns the length of the given vector.
func (v Vec2[T]) Len() T {
	return T(math.Sqrt(float64(v.Dot(v))))
}

// Unit computes the unit vector along the direction of the given vector.
func (v Vec2[T]) Unit() Vec2[T] {
	n := 1.0 / v.Len()
	return Vec2[T]{v.X * n, v.Y * n}
}
