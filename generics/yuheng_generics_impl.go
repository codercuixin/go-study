package main
func test_imp[T any](x T){
	println("x:", x)
}

func main(){
	test_imp(1)
	type X int 
	// test underlying type
	test_imp(X(2))
	test_imp("1")
}