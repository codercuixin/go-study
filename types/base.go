package main

import (
	"fmt"
	"math"
	"reflect"
	"sync/atomic"
)

func main(){
	// testDiffBaseNum()
	// a := 111_22_33_44 
	// println(a)

	//use of untyped nil in assignment
	// x := nil 

	// var m map[string]int = nil 
	// var c []int = nil 
	// //invalid operation: m == c (mismatched types map[string]int and []int)
	// fmt.Println(m == nil, c==nil)
	// atomictest()
	testAlias()
}
//就是语法糖，编译器会将别名还原为原始类型
func testAlias(){
	type X = int 
	var a int = 100
	var b X = a 
	fmt.Printf("%T, %v\n", b, b)
	fmt.Println(reflect.TypeOf(b))
}

func atomictest(){
	var x atomic.Int64 
	go func(){
		x.Add(1)
	}()
	x.Store(x.Load() + 1 )
	println(x.Load())
	
}


func testDiffBaseNum(){
	a, b, c := 0b1010, 0o144, 0x64

	fmt.Printf("0b%b, %#o, %#x\n", a, b, c)
	fmt.Println(math.MinInt8, math.MaxInt8)
}