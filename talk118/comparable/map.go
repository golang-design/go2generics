package main

type P map[interface{}]struct{}     // OK
type R[T comparable] map[T]struct{} // OK
// type Q[T any] map[T]struct{}     // ERROR: incomparable map key type T
