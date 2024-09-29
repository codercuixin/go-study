package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.Write([]byte("abcd"))
	br := make([]byte, 1, 1)
	n, err :=buf.Read(br);
	if err !=nil{
		panic(err)
	}
	fmt.Println("read ", n, " bytes from buffer")
	fmt.Println(string(br))
	fmt.Println("buf left is: ", buf.String())

}