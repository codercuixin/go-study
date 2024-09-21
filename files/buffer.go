package main

import "bytes"
import "fmt"

func main() {
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte{1, 2, 3})
	fmt.Println(buf.Len())
}