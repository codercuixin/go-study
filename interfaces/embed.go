package main

import (
	"fmt"
	"reflect"
)

type Aer interface {
	Test()
}

type Ber interface {
	ToString(string) string
}

type Xer interface {
	Aer
	Ber
	ToString(s string) string //签名相同，并集去重
	// ToString()string //不允许重载
	Print()
}

type N struct{}

func (N) ToString(string) string {
	return ""
}
func (N) Test() {

}
func (N) Print() {

}

func main() {
	i := Xer(N{})
	t := reflect.TypeOf(i)
	
	for i:=0; i<t.NumMethod(); i++{
		fmt.Println(t.Method(i))
	}
}
