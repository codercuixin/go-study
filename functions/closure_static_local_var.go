package main 

func testStaticLocalVar() func(){
	n := 0 
	return func() {
		n++
		println("call:", n)
	}
}

func main(){
	f := testStaticLocalVar()
	f()
	f()
	f()
}