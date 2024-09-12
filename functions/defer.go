package main

import (
	"io"
	"log"
	"os"
)

func main(){
	f, err := os.Open("./defer.go")
	if err != nil{
		log.Fatalln(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil{
		log.Fatalln(err)
	}
	println(string(b))
}