package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

func main() {
	var x int64 = 0x11223344

	big := bytes.NewBuffer(nil)
	binary.Write(big, binary.BigEndian, x)

	litte := bytes.NewBuffer(nil)
	binary.Write(litte, binary.LittleEndian, x)

	fmt.Printf("B: %x\n", big.Bytes())
	fmt.Printf("L: %x\n", litte.Bytes())

	// 判断系统是否小端
	y :=1 
	p := (*byte)(unsafe.Pointer(&y))
	println(*p ==1)
}

