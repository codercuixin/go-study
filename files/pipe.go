package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r, w := io.Pipe()
	go func(){
		defer w.Close()
		fmt.Fprintf(w, "hello, world!\n")
	}()
	io.Copy(os.Stdout, r)
}