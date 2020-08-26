// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notsupport

// variadic type parameter list is not supported.
type Elements[Ts ...comparable] struct {
	es ...Ts
}

type Metric[Ts ...comparable] struct {
	mu sync.Mutex
	m  map[Elements[Ts...]]int
}

func (m *Metric[Ts...]) Add(vs ...Ts) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.m == nil {
		m.m = make(map[Elements[Ts...]]int)
	}
	m[Elements[Ts...]{vs...}]++
}