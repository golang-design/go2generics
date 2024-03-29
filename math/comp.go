// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Ordered is a type constraint that matches all ordered types.
// (An ordered type is one that supports the < <= >= > operators.)
// In practice this type constraint would likely be defined in
// a standard library package.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 | ~string
}

// Min returns the minimum among all parameters
func Min[T Ordered](s ...T) T {
	if len(s) == 0 {
		panic("Min of no elements")
	}
	r := s[0]
	for _, v := range s[1:] {
		if v < r {
			r = v
		}
	}
	return r
}

// Max returns the maximum among all parameters
func Max[T Ordered](v0 T, vn ...T) T {
	switch l := len(vn); {
	case l == 0:
		return v0
	case l == 1:
		if v0 > vn[0] { return v0 }
		return vn[0]
	default:
		vv := Max(vn[0], vn[1:]...)
		if v0 > vv { return v0 }
		return vv
	}
}