package main 

import "fmt"

type A int 

type B string 

func (a A) Test(){
	println("A:", a)
}
func(b B) Test(){
	println("B:", b)
}

type Tester interface{
	A | B 
	Test()
}

func test[T Tester](v T){
	v.Test()
}

func main(){
	test[A](1);
	test[B]("1")
	test2(1)
	fmt.Println(makeSlice(10))
	testArray([]int{1,2,3})
	testArray([]float64{1.1,2.2,3.3})
}


func test2[T int| float32] (x T){
	fmt.Printf("%T, %v\n", x, x)
}


func makeSlice[T int| float64](x T) []T{
	s := make([]T, 10)
	for i:=0; i<cap(s); i++{
		s[i] = x
	}
	return s
}

func testArray[T int|float64, E ~[]T](x E){
	for _, v := range x{
		fmt.Println(v)
	}
}


type Num int 
func (n Num) print(){
	println("Num:", n)
}
func testNum[T Num](n T){
	// n.pr÷int()
}

func testSwitch[T any](x T){
	switch any(x).(type){
	case Num:
		println("Num")
	case int:
		println("int")
	case float64:
		println("float64")
	}
}