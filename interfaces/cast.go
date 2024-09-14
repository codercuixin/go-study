package main

import (
	"fmt"
	"reflect"
)
type Aer interface{
	toString(string)string 
}

type Xer interface{
	test()
	toString(string)string 
}

type N struct{}
func (*N) test(){

}
func (N) toString(s string)  string{
	return s 
}

func main(){
	// basic()
	// typeInfer()
	// switchInfer()
	// cmp()
	// isNil()
	isNil2()
}



func isNil(){
	var x Xer = (*N)(nil)  //type !=nil and val is nil
	println(x == nil)
	println(x == nil || reflect.ValueOf(x).IsNil())
}

func isNil2(){
	var t1 any 
	var t2 any = ([]int)(nil)

	println(t1 ==nil)
	println(t2 ==nil)
	
	println(t1 ==nil || reflect.ValueOf(t1).IsNil())
	println(t2 == nil || reflect.ValueOf(t2).IsNil())
}

//use runtime.efaceeq, runtime.ifaceeq; 
func cmp(){
	println(any(nil) == any(nil))
	println(any(100) == any(100))
}



func switchInfer(){
	var i any = &N{}
	switch v := i.(type){
	case nil:
	case *int:
	case func()string:
	case *N: fmt.Println("*N, ", v)
	case Xer: fmt.Println(v)
	default:
	}
}

func typeInfer(){
	var x Xer = &N{} //super
	var a Aer = x // sub

	n, ok := a.(*N)
	fmt.Println(n, ok) //true

	x2, ok := a.(Xer)
	fmt.Println(x2, ok)

	var e any = (*int)(nil) // with type nil
	_, ok = e.(*int)
	println(ok)
}


func basic(){
	var x Xer = &N{}
	var a Aer = x 

	a.toString("abc")
	//cannot use a (variable of type Aer) as Xer value in variable declaration:
	// Aer does not implement Xer (missing method test)
	// var x2 Xer = a 
	// x2 := Xer(a)
}
