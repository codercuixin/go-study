package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	data := []byte{0, 1, 2, 3, 4, 5, 6}
	r := bytes.NewReader(data)
	limit_reader := io.LimitReader(r, 3)
	for{
		buf := make([]byte, 2)
		n, err := limit_reader.Read(buf)
		fmt.Println(n, err, buf)
		if err == io.EOF{
			break
		}
	}
}