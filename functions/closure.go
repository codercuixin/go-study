package main 
func closure() func(){
	x := 100
	println(&x, x )
	return func() {
		x ++
		println(&x, x)
	}
}

func main(){
	closure()()
}


