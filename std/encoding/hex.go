package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	b := []byte{0x11, 0x22, 0x33}
	s := hex.EncodeToString(b)
	b2, err := hex.DecodeString(s)
	if err != nil {
		log.Fatalln(err)
	}
	if !bytes.Equal(b, b2) {
		log.Fatal()
	}
	fmt.Printf("%s\n% x\n", s, b2)

	b = []byte("hello world")
	fmt.Println(hex.Dump(b))
}
