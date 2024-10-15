package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//  copy()
	//  copyBuffer()
	copyN()
}
func copy(){
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil{
		log.Fatal(err)
	}
}

func copyN(){
	r := strings.NewReader("some io.Reader stream to be read\n")
	if n, err := io.CopyN(os.Stdout, r, 100); err!=nil{
		log.Fatalf("written is %v, error is %v", n, err)
	}
}

func copyBuffer(){
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err !=nil{
		log.Fatal(err)
	}
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil{
		log.Fatal(err)
	}
}