package main 
type Aer interface{
	test()
}
type Ber interface{
	test()
}
type Cer interface{
	toString()string 
}
type Xer interface{
	test()
	toString() string 
}
type N struct {
	x int 
}
func (N) test(){

}
func (N) toString()string{
	return ""
}
func main(){
	a, b := N{100}, N{100}
	// println(a == b)
	// println(Aer(a) == Aer(b))
	// println(Aer(&a) == Aer(&b))

	// println(Aer(a) == Ber(b))
	// println(Aer(a)  == Xer(b))
//cannot convert b (variable of type N) to type Cer: 
//N does not implement Cer (wrong type for method toString)
	// println(Aer(a) == Cer(b))

	println(Aer(a) == b)
}

