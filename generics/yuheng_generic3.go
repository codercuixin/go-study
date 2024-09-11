package main

import "fmt"

type N struct{}

func (N) String() string {
	return "N"
}

func test[T fmt.Stringer](v T) {
	fmt.Println(v)
}


func main(){
	test(N{})
}