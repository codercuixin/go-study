package main

import (
	"errors"
	"io/fs"
	"os"
	"fmt"
)

func main(){
	// simple()
	custom()
}

func simple(){
	if _, err := os.Open("non-existing"); err !=nil{
		var pathError *fs.PathError
		if errors.As(err, &pathError){
			fmt.Println("Failed at path:", pathError.Path)
		}else{
			fmt.Println(err)
		}
	}
}

func custom(){
	x := &TestError{}
	y := fmt.Errorf("y, %w", x)
	z := fmt.Errorf("z, %w", y)
	//wrap error
	var x2 *TestError
	if errors.As(z, &x2){
		fmt.Println(x == x2)
	}

}

type TestError struct{

}
func(*TestError) Error() string{
	return ""
}