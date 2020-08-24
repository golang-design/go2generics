// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

import "testing"

func TestMetric1(t *testing.T) {
	m := Metric1[int]{}

	m.Add(1)
	m.Add(2)
}

func TestMetric2(t *testing.T) {
	m := Metric2[int, string]{}

	m.Add(1, "1")
	m.Add(2, "2")
}

func TestMetric3(t *testing.T) {
	m := Metric3[int, string, float64]{}

	m.Add(1, "1", 1.0)
	m.Add(2, "2", 2.0)
}