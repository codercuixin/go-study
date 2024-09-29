package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	fmt.Printf("%d\n", b.Len())
	os.Stdout.Write(b.Bytes())
	
	fmt.Println(b)
	b.Write([]byte("A"))
	fmt.Printf("%d\n", b.Len())
}
