package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("err1")
	const name, id = "bimmler", 17
	err2 := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err1)
	fmt.Println(err2)
}