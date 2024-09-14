package main

import "fmt"

type FuncString func() string

func (f FuncString) String() string {
	return f()
}

func main() {
	f := func() string {
		return "hello world"
	}
	var t fmt.Stringer = FuncString(f)
	//fmt.Println(t)
	fmt.Println(t.String())
}