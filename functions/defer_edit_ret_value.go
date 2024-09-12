package main 

func deferEditNamedRetValue()(z int){
	defer func(){
		z += 200
	}()
	return 100
}

func deferEditLoalValue()(int){
	z :=0
	defer func(){
		z += 200
	}()
	z = 100
	return z
}



func main(){
	println(deferEditNamedRetValue())
	println(deferEditLoalValue())
}
