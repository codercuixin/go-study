package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r, w := io.Pipe()
	go func(){
		fmt.Fprintf(w, "some io.Reader tream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil{
		log.Fatal(err)
	}
	

}