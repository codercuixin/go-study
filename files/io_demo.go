package main

import (
	"bufio"
	"log"
	"os"
)

func main(){
	f, err := os.Create("./demo.txt")
	if err != nil{
		log.Fatalln(err)
	}
	defer f.Close()
	defer f.Sync()
	w := bufio.NewWriter(f)
	defer w.Flush()
	w.WriteString("hello world")
}