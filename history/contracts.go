// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package history

// generics: by contract

contract Bigger(T) {
    IsBiggerThan(T) bool
}

func Max(type T Bigger) (a, b T) T { ... }

// generics: by parametric interface

type Bigger(type T) {
    IsBiggerThan(T)
}

func Max(type T Bigger(T)) (a, b T) T { ... }

// a more complex example:

contract C(P1, P2) {
    P1 m1(x P1)
    P2 m2(x P1) P2
    P2 int, float64
}

func F(type P1, P2 C) (x P1, y P2) P2 { ... }

type I1 (type P1) interface {
    m1(x P1)
}
type I2 (type P1, P2) interface {
    m2(x P1) P2
    type int, float64
}

func F(type P1 I1(P1), P2 I2(P1, P2)) (x P1, y P2) P2 { ... }