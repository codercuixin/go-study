package main

import "fmt"

type Data[T any] struct {
	x T
}

func (d Data[T]) test() {
	fmt.Println(d)
}

// func (d Data[T]) test2[X any(x X) {
// 	fmt.Println(d)
// }

func main() {
	d :=Data[int]{x : 1}
	d.test()
}