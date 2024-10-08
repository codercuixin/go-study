package main

import "reflect"
import "fmt"

func main() {
	// testOf()
	// testValAndPinter()
testElem()
}

func testOf(){
	tByte := reflect.TypeOf(byte(0))
	tInt := reflect.TypeOf(0)
	tString := reflect.TypeOf("")
	tSlice := reflect.TypeOf([]int{})

	a := reflect.ArrayOf(10, tByte) // [10]uint8
	c := reflect.ChanOf(reflect.BothDir, tInt) //chat int 
	m := reflect.MapOf(tString, tInt) // map[string]int
	s := reflect.SliceOf(tInt) // []int
	p := reflect.PointerTo(tInt) //*int

	in := []reflect.Type{tInt, tInt, tSlice}
	out := []reflect.Type{tString}
	f := reflect.FuncOf(in, out, true) // func(int, int, ...int) string
	fmt.Println(a,c, m, s, p, f)
}

func testValAndPinter(){
	x := 100 
	t, p := reflect.TypeOf(x), reflect.TypeOf(&x)
	fmt.Println(t, p, t == p) //int *int false
	fmt.Println(t.Kind(), p.Kind())
	// 方法 Elem, 返回Array, Chan, Map, Pointer, or Slice 的基类型
	fmt.Println(t == p.Elem())
}


func testElem(){
	x :=  100
	p := reflect.TypeOf(&x)
	s := reflect.TypeOf([]int32{})
	m := reflect.TypeOf(map[string]int{})

	fmt.Println(p.Elem())
	fmt.Println(s.Elem())

	fmt.Println(m.Key())
	fmt.Println(m.Elem())
}