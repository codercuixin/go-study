package main

import "fmt"
import "io"

type MyInterface interface {
	M1(int) error
	M2(io.Writer, ...string)
}

type Interface1 interface {
	M1()
	test()
}

type Interface2 interface {
	M1()
	M2()
}

type Interface3 interface {
	Interface1
	Interface2
}

type EmptyInterface interface {
}

type C1 interface {
	~int | ~int32
	M1()
}

type T struct {
}

func (T) M1() {}

type T1 int

func (T1) M1() {

}

func foo[P C1](t P) {

}

func Sort[Elem interface {
	Less(y Elem) bool
}](list []Elem) {

}

type book struct{}

func (x book) Less(y book) bool {
	return true
}

func main() {
	fmt.Println("Hello, Go!")
	var t1 T1
	foo(t1)

	var bookshelf []book
	Sort[book](bookshelf)

	var iv = Vector[int]{}
	var sv Vector[string]
	sv = []string{"a", "b", "c", "d"}
	iv.Dump()
	sv.Dump()
}

type Vector[T any] []T

func (v Vector[T]) Dump() {
	fmt.Printf("%#v\n", v)
}
