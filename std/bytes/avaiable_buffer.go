package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var buf bytes.Buffer
	fmt.Println(buf.Available())
	for i:=0; i<4; i++{
		ab := buf.AvailableBuffer()
		ab = strconv.AppendInt(ab, int64(i), 10)
		ab = append(ab, ' ')
		buf.Write(ab)
	}
	fmt.Println(buf.Available())
	os.Stdout.Write(buf.Bytes())

}