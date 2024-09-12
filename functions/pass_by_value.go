package main 

//go:noline 
func test(x *int, s[]int){
	println(*x, len(s))
}

func main(){
	x := 100 
	s := []int{1,2,3}
	test(&x, s)

	var p *int 
	test2(&p)
	println(*p)
}


func test2(p **int){
	x := 100 
	*p = &x 
}

