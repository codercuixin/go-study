package main 
import "reflect"
import "fmt"
type E int 
func (e E) V(){}
func (e *E) P(){}

func show(i interface{}){
	t := reflect.TypeOf(i)
	for i:=0; i< t.NumMethod(); i++{
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
func main(){
	println("T{E}")
	show(struct{E}{})
	println("T{*E}")
	show(struct{*E}{})
	println("*T{E}")
	show(&struct{E}{})
	println("*T{*E}")
	show(&struct{*E}{})
}