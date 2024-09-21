package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	r1 := bytes.NewReader([]byte{0, 1, 2, 3, 4, 5, 6})
	r2 := bytes.NewReader([]byte{10 ,11, 12,13})

	w1 := bytes.NewBuffer(nil)
	w2 := bytes.NewBuffer(nil)
	//read from r1 ,then r2
	mr := io.MultiReader(r1, r2)
	//write to w1 &w2, at the same time
	mw := io.MultiWriter(w1, w2)
	for{
		buf :=make([]byte, 2)
		n, err := mr.Read(buf)
		fmt.Println(n, err, buf)

		if err != nil{
			break
		}
		_, _ = mw.Write(buf[:n])
	}
	fmt.Println(w1.Bytes())
	fmt.Println(w2.Bytes())

}