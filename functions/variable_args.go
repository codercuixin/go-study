package main 
import "fmt"

func testVariableArgs(s string, a ...int){
	fmt.Printf("%T, %v\n", a,a)
}

func testChangeSelf( a ...int){
	 for i := range a{
		a[i] += 100
	 }
}

func main(){
	testVariableArgs("abc", 1, 2, 3, 4)

	a := [3]int{10 ,20 ,30}
	testVariableArgs("test", a[:]...)

	testChangeSelf(a[:]...)
	fmt.Println(a)
}