package main 

type E struct{}
type T struct{
	E
}

func (E) toString() string{
	return "E"
}

func (T) toString(s string) string{
	return "T" + s
}



func main(){
	var t T
	//yh_anno_nest2.go:20:10: not enough arguments in call to t.toString
	println(t.toString())
	println(t.E.toString())
}