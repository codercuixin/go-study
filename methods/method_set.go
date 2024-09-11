package main 
import (
	"fmt"
	"reflect"
)
type T int 
func (T) A(){}
func (T) B(){}
func (*T) C(){}
func (*T) D(){}

func show(i interface{}){
	t := reflect.TypeOf(i)
	for i:=0; i< t.NumMethod(); i++{
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

type Xer interface {
	B()
	C()
}


func main(){
	// var t T
	// var pt *T
	// show(t)
	// println("---")
	// show(pt)

	var n T =1 
	n.B()
	n.C()

	var x Xer = &n
	x.B()
	x.C()
}

