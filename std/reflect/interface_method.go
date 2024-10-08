package main

import "reflect"
import "fmt"
import "unsafe"

func main() {
	// testConvertibleTo()
	// testSubType()
	testCanSet()
}

func testConvertibleTo() {
	a := 100
	v := reflect.ValueOf(&a)

	if t := reflect.TypeOf((*int)(nil)); v.Type().ConvertibleTo(t) {
		if p, ok := v.Interface().(*int); ok {
			*p += 100
			fmt.Println(*p, a)
		}
	}
}

type X int

func (X) String() string {
	return ""
}

func testSubType() {
	var x X
	t := reflect.TypeOf(x)

	st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(t.Implements(st))
	fmt.Println(t.AssignableTo(st))

}


func testCanSet(){
	a:=100
	v := reflect.ValueOf(a)
	p := reflect.ValueOf(&a)
	//A value is addressable if it is an element of a slice, an element of an addressable array, a field of an addressable struct,
	/// or the result of dereferencing a pointer
	fmt.Println(v.CanAddr(), v.CanSet())
	fmt.Println(p.CanAddr(), p.CanSet())

	// *data
	e := p.Elem()
	fmt.Println(e.CanAddr(), e.CanSet())
	if e.CanSet(){
		e.SetInt(200)
		fmt.Println(a)
	}

	if e.CanAddr(){
		*(*int)(unsafe.Pointer(e.UnsafeAddr())) = 300
		fmt.Println(a)
	}
}