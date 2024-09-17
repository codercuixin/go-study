package main

import "fmt"
import "testing"
import "os"

func addXY(x, y int)int{
	return x + y 
}
func ExampleAdd() {
	fmt.Println(addXY(1, 2))
	fmt.Println(addXY(2, 2))

	//Output:
	//3
	//3
}

func TestMain(m *testing.M){
	//setup
	println("setup")
	code := m.Run()
	//teardown
	println("teardown")
	os.Exit(code)
}