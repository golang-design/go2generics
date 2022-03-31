package main

import "fmt"

// type Stringish interface {
// 	~string | fmt.Stringer
// }

type Stringish interface {
	~string
}

func ToString[T Stringish](s T) string {
	switch x := any(s).(type) {
	case string:
		return string(x)
	case fmt.Stringer:
		return x.String()
	default:
		panic("impossible")
	}
}

type T string

func main() {
	ToString(T("x")) // panic
}
