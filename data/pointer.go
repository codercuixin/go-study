package main

import "unsafe"

func main(){
	// basic()
	// sizeOfNilPointer()
	equal()
}

func equal(){
	var p1, p2 *int 

	println(p1 == p2)
	var x int 
	p1, p2 =&x, &x 
	println(p1 == p2)

	var y int 
	p2 = &y 
	println(p1 == p2 )
}

func sizeOfNilPointer(){
	var p *int 
	println(unsafe.Sizeof(p))

	var x  int 
	var pp **int = &p 
	*pp = &x 

	*p =100
	**pp +=1 
	println(**pp, *p, x)

}
func basic(){
	var x int 
	var p *int = &x 

	*p = 100
	println(p, *p, &x)
}