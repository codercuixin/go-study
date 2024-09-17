package main 

import (
	"go-study/packages/mylib"
	_ "go-study/packages/mylib/x"
	_ "go-study/packages/mylib/x/y"
	//could not import go-study/packages/mylib/internal/a (invalid use of internal package "go-study/packages/mylib/internal/a
	// _ "go-study/packages/mylib/internal/a"
)
// import "unsafe"
import "fmt"

func main(){
	//undefined:
	// z := mylib.add(1, 2)
	 
	mylib.Hello()
	
	p := mylib.NewData()
	//p.x undefined (cannot refer to unexported field 
	// p.x =100
	fmt.Println(p)

	// d := (*struct{
	// 	_ int 
	// 	y int 
	// })(unsafe.Pointer(p))
	// d.y = 100
	// fmt.Println(*d)

	// d := mylib.DataTmp{}
	// d.SetY(100)
	// d.Test()
}