package main 
type N int 
func (n N) copy(){
	println("copy.value is", n)
}
func (n *N) ref(){
	println("ref.value is", *n)
}
func main(){
	// var n N = 1
	// // expression 
	// var e func(*N) () = (*N).ref 
	// e(&n)

	// //value 
	// var v func() = n.ref 
	// v()

	var n N =100
	var v func() = n.copy
	n++ 
	n.copy() 

	v()
	test(v)

	
	
}
func test(v func()){
	v()
}