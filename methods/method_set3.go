package main 
import "reflect"
type X int 
func (*X) A(){
	println("X.A")
}
type Y = X 
func(*Y) B(){
	println("Y.B")
}

func main(){
	var x X 
	x.A()
	x.B()
	println("---")
	var y *Y = nil
	y.A()
	y.B()
	
	t := reflect.TypeOf(&x)
	println(t.NumMethod())
	for i, n := 0, t.NumMethod(); i<n; i++{
		m := t.Method(i)
		println(m.Name, m.Type)
	}
}