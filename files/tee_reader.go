package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	data := []byte{0, 1, 2, 3, 4, 5, 6}
	r := bytes.NewReader(data)
	w := bytes.NewBuffer(make([]byte, 0, 7))
	t := io.TeeReader(r, w)
	for{
		buf := make([]byte, 2)
		n, err := t.Read(buf)
		fmt.Println(n, err, buf)
		if err == io.EOF{
			break
		}
	}
	fmt.Println(w.Bytes())
}