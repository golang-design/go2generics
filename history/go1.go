// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package history

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
func MaxUintptr(a, b uintptr) uintptr {
	if a > b {
		return a
	}
	return b
}
