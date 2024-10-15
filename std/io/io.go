package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"os"
)

func main() {
	// readAll()
	// readAtLeast()
	// readFull()
	// limitReader()
	// multiReader()
	// teeReader()
	// sectionReader()
	multiWriter()
}

func multiWriter(){
	r := strings.NewReader("some io.Reader stream to be read\n")
	var buf1, buf2 strings.Builder
	w := io.MultiWriter(&buf1, &buf2)

	if _, err := io.Copy(w, r); err !=nil{
		log.Fatal(err)
	}
	fmt.Print(buf1.String())
	fmt.Print(buf2.String())
}

func sectionReader(){
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)
	fmt.Println(s.Size())
	// if _, err := io.Copy(os.Stdout, s); err != nil{
	// 	log.Fatal(err)
	// }

	// buf := make([]byte, 9)
	// if _, err := s.Read(buf) ;err != nil{
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", buf)

	// buf := make([]byte, 6)
	// if _, err := s.ReadAt(buf, 10); err != nil{
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", buf)

	if _, err := s.Seek(10, io.SeekStart); err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}

}

func teeReader(){
	var r io.Reader = strings.NewReader("some io.Reader stream to be read\n")
	// everyting reader from r will write to os.Stdout
	r = io.TeeReader(r, os.Stdout)
	if b, err := io.ReadAll(r); err != nil{
		log.Fatal(err)
	}else{
		fmt.Printf("\nbuf: %s\n", b)
	}
}

func multiReader(){
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func limitReader(){
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}

}

func readAll() {
	r := strings.NewReader("Go is a general-purpose language designed with systme programming in mind.")
	b, err := io.ReadAll(r)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
	b, err = io.ReadAll(r)
	fmt.Printf("%s, %v", b, err)
}

func readAtLeast(){
	r := strings.NewReader("some io.Reader stream to be read\n")
	buf := make([]byte, 14)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buf)

	//buffer smaller than minimal read size
	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err !=nil{
		fmt.Println(err)
	}

	longBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, longBuf, 64); err !=nil{
		fmt.Println(err)
	}
	
}

func readFull(){
	r := strings.NewReader("some io.Reader stream to be read\n")
	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, r.Len()+1)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s\n", longBuf)

}