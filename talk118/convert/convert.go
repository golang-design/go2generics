package main

func Foo[T int]() T {
	x := 42
	return T(x) // OK
	// return int(x) // ERROR: cannot use int(x) (value of type int) as type T in return statement
}
