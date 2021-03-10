package fmt

import "fmt"

// Print prints the elements of a slice.
// It should be possible to call this with any slice value.
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func Pirnt[T any](s ...T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
