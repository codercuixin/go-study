package main

import "reflect"

func main(){
	// type Xer interface{
	// 	A()
	// 	B()
	// }
	// var i Xer 
	
	// t := reflect.TypeOf(i)
	// println(t !=nil)

	// v := reflect.ValueOf(i)
	// println(v.IsValid())
	// //panic: reflect: call of reflect.Value.Type on zero Value
	// // println(v.Type())

	// t = reflect.TypeOf((*Xer)(nil))
	// println(t != nil)
	// println(t.Elem().Name())

	// test2()
	test3()
	
}

func test2(){
	var i interface{} = (*int)(nil)
	t := reflect.TypeOf(i)
	println(t !=nil)

	v := reflect.ValueOf(i)
	println(v.IsValid())
	println(v.Type())
}

func test3(){
	//zero value
	var v reflect.Value 
	println(v.IsValid())

	//panic: reflect: call of reflect.Value.IsZero on zero Value
	// println(v.IsZero())

	//panic: reflect: call of reflect.Value.IsNil on zero Value
	println(v.IsNil())


	v = reflect.ValueOf([]int{})
	println(v.IsValid())
	println(v.IsZero())

	var i interface{} = (*int)(nil)
	v = reflect.ValueOf(i)
	println(v.IsValid())
	println(v.IsNil())
}