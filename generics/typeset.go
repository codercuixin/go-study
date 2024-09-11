package main 
type Intf1 interface {
    ~int | string
  F1()
  F2()
}

type Intf2 interface {
  ~int | ~float64
}

type I interface {
    Intf1 
    M1()
    M2()
    int | ~string | Intf2
}

func doSomething[T I](t T){

}
type MyInt int 

func (MyInt) F1(){

}
func (MyInt) F2(){
	
}
func (MyInt) M1(){
	
}
func (MyInt) M2(){
	
}

type MyString string 

func (MyString) F1(){

}
func (MyString) F2(){
	
}
func (MyString) M1(){
	
}
func (MyString) M2(){
	
}

func main(){
	var a int = 11
	// doSomething(a)

	var b = MyInt(a)
	doSomething(b)
	// 
	var s = MyString("hello")
	doSomething(s)
}