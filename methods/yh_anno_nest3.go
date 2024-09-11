package main


type E1 struct {
	x int
}
type T1 struct {
	E1
	x string 
}

func (e *E1) do(){
	e.x = 100
	println(e.x)
}

func (t *T1) do(){
	t.x = "hello"
	println(t.x)
}

func main(){
	var t T1
	t.do()
	t.E1.do()
}