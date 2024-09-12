package main 

import "unsafe"

func main(){
	testNew()
}

func testNew(){
	s := *new([]int) // slice:{ptr, len, cap}
	m := *new(map[string]int) //8
	c := *new(chan int) // 8

	println(unsafe.Sizeof(s))
	println(unsafe.Sizeof(m))
	println(unsafe.Sizeof(c))
}


func testMake(){
	s := make([]int, 10, 100)
	println(unsafe.Sizeof(s))

	m := make(map[string]int, 10)
	m["a"]=1
	println(m)
}