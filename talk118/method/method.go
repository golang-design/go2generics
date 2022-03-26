package main

type X struct{}

func (x X) Foo[T any](t T) {}

func Bar(x X) {
	x.(interface{ Foo(int) }) // compiler cannot infer to generate　X.Foo[int]
}
