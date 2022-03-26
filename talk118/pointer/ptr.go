package main

import "fmt"

type A struct{}

func (a A) Foo() {}
func (a A) Bar() {}

type B struct{}

type C[T any] interface {
	*T

	Foo()
	Bar()
}

func (b B) Foo()  {}
func (b *B) Bar() {}

func Want[U any, V C[U]]() (x V) { // Guarantee U must be a pointer.
	return
}

func main() {
	a := Want[A]()
	b := Want[B]()

	fmt.Printf("%T, %T\n", a, b) // *main.A, *main.B
}
