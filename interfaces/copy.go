package main

import "strconv"

type Xer1 interface {
	toString() string
}

type N1 struct {
	x int
}

func (n N1) toString() string {
	return strconv.Itoa(n.x)
}

func main(){
	//结构体会复制目标对象，通常以指针替代原始值。
	n := N1{100}
	var x Xer1 = n 
	n.x = 200

	println(n.toString())
	println(x.toString())
}