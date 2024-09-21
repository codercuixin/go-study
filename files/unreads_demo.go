package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []byte{0, 1, 2, 3, 4, 5, 6}
	r := bytes.NewReader(data)

	b, err := r.ReadByte()
	fmt.Println(b, err)
	//须在 ReadByte 之后，位置 seek-1
	err = r.UnreadByte()
	if err !=nil{
		fmt.Println("unread:", err)
	}

	for{
		b, err := r.ReadByte()
		fmt.Println("for: ", b, err)
		if err != nil{
			break
		}
	}
	
}