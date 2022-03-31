package main

import "math"

type Vec2[T ~float32 | ~float64] struct {
	X, Y T
}

func NewV[T ~float32 | ~float64](x, y T) Vec2[T] {
	return Vec2[T]{x, y}
}
func main() {
	v := NewV(math.MaxFloat32, math.MaxFloat32)
	switch (any)(v).(type) {
	case Vec2[float32]:
	case Vec2[float64]:
		panic("?") // 会执行 panic
	}
}
