package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	var s fs.FS = os.DirFS("/dev")
	f, err := s.Open("urandom")
	if err != nil{
		log.Fatalln(err)
	}
	defer f.Close()

	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Printf("%#v, %d\n", buf, n)
}