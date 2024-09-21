package main

import (
	"bytes"
	"io"
	"fmt"
)

func main() {
	data := []byte{0, 1, 2, 3, 4, 5, 6}
	r := bytes.NewReader(data)

	for i:=0; i<3; i++{
		if i == 2{
			r.Seek(-2, io.SeekEnd)
		}
		buf := make([]byte, 1)
		n, err := r.Read(buf)
		fmt.Println("read:", n, err, buf)
	}

	//从头开始，不影响 Seek
	buf := make([]byte, 1)
	n, err := r.ReadAt(buf, 1)
	fmt.Println("ReadAt:", n, err, buf)

	//继续 for.Read 位置
	b, err := r.ReadByte()
	fmt.Println("ReadByte:", b, err)
}