package main 

type  Xer interface{
	test()
	toString()string 
}
type N struct{}
func (*N) test(){

}

func(N) toString()string {
	return ""
}

func main(){
	var n N 
	//cannot use n (variable of type N) 
	//as Xer value in variable declaration: N does not implement Xer (method test has pointer 
	// var t1 Xer = n 
	var t Xer = &n 
	t.test()
	t.toString()
}

