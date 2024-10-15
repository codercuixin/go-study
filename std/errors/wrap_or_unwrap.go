package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("error1")
	err2 := fmt.Errorf("error2: [%w]", err1)
	// err3 := fmt.Errorf("error3: [%w]", err2)
	fmt.Println(err2)
	fmt.Println(errors.Unwrap(err2))
	fmt.Println(errors.Unwrap(err1))

	fmt.Println("err2 is err1: ", errors.Is(err2, err1))
}