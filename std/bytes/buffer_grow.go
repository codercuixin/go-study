package main

import "bytes"
import "fmt"

func main() {
	var buf bytes.Buffer
	buf.Grow(64)
	buf.Write([]byte("hello world"))
	bb := buf.Bytes()
	fmt.Printf("%q",bb[:buf.Len()])
}