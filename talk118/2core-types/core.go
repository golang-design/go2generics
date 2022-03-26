package main

type X float32 // The core type of X is float32
type Y float32 // The core type of Y is float32

type U interface { // The core type of U is *int
	*int
	String() string
}

type V interface { // The core type of V is float32
	~float32
	String() string
}

type W interface { // W has no core type.
	int | float32
}
