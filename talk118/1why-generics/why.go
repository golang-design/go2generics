package main

type T interface {
	Add(T)
}

func Sum(elems ...T) (sum T) { // T 可以是任何实现了 Add() 方法的类型
	for i := range elems {
		sum.Add(elems[i])
	}
	return
}

func GenericSum[T ~int](elems ...T) (sum T) { // T 的底层类型必须底层类型为 int 约束的类型
	for i := range elems {
		sum += elems[i]
	}
	return
}
