package main

import (
	"fmt"
	"reflect"
)

type N int

func (N) X() {

}

func (*N) Y() {

}

type M struct {
	N
}

func (M) A() {

}

func (*M) b() {

}

func (*M) C() {

}
func main() {
	// printStructMethodSet()
	printInterfaceMethodSet()
}

func printInterfaceMethodSet(){
	// with unexported method in interface
	var i interface {
		A()
		b()
		C()
	}
	// t := reflect.TypeOf(i)
	t := reflect.TypeOf(&i).Elem()
	

	for i:=0; i<t.NumMethod(); i++{
		fmt.Println(t.Method(i))
	}
}


func printStructMethodSet(){
	t := reflect.TypeOf((*M)(nil))
	for i:= 0; i<t.NumMethod(); i++{
		fmt.Println(t.Method(i))
	}

	if m, ok := t.MethodByName("C"); ok{
		fmt.Println(m)
	}
}