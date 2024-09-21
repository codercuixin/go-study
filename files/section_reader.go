package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	data := []byte{0, 1, 2, 3, 4, 5, 6}
	r := bytes.NewReader(data)
	section_reader := io.NewSectionReader(r, 1, 5)
	for{
		buf := make([]byte, 2)
		n, err := section_reader.Read(buf)
		fmt.Println(n, err, buf)
		if err == io.EOF{
			break
		}
	}
}