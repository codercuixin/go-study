package main 

type E struct{}
type T struct{
	E
}

func (E) toString() string{
	return "E"
}

func (T) toString() string{
	return "T"
}



func main(){
	var t T
	println(t.toString())
	println(t.E.toString())
}