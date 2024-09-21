package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	f, _ := os.Open("./bufio.go")
	defer f.Close()
	var r io.Reader = f 
	for{
		buf := make([]byte, 512)
		n, err := r.Read(buf)

		if err == io.EOF{
			break
		}
		fmt.Printf("read %d bytes\n", n)
	}
}