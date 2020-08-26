// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constraints

// An interface specifies methods and types for each of the
// type parameters it constrains.
type Stringer interface {
	String() string
}

type PrintableStringer[T Stringer] interface {
	Print()
}

type Sequence interface {
	type string, []byte
}
