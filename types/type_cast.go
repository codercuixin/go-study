package main

import "fmt"

func main() {
	complex()
}

func complex() {
	x := 100
	p := (*int)(&x)
	fmt.Println(*p)
}

func simple() {
	a := 10
	b := byte(a)
	c := a + int(b)
	println(c)
}