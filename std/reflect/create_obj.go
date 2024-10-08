package main 

import . "reflect"
import "fmt"
import "unsafe"

func main(){
	v := New(TypeOf(0))
	v.Elem().SetInt(100)
	fmt.Println(v.Elem().Int())

	p := new(int)
	v = NewAt(TypeOf(0), unsafe.Pointer(p))
	v.Elem().SetInt(200)
	fmt.Println(v.Elem().Int(), *p)

	v = MakeSlice(TypeOf([]int{}), 0, 10)
	v = Append(v, ValueOf(1), ValueOf(2))
	fmt.Println(v)

}