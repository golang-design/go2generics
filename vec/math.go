// Copyright 2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vec

import "math"

// ApproxEq approximately compares v1 and v2.
func ApproxEq[T floats](v1, v2, epsilon T) bool {
	return math.Abs(float64(v1)-float64(v2)) <= float64(epsilon)
}
