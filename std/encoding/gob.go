package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func main() {
	a := struct {
		X int
		Y string
		z int16 // ignore
	}{

		100,
		"abc",
		12,
	}

	b := struct {
		B bool
		Y *string
		X *int32
		Z int16
	}{}
	data := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(data)
	if err := enc.Encode(&a); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("% x\n", data.Bytes())
	dec := gob.NewDecoder(data)
	if err := dec.Decode(&b); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", b)
	fmt.Println(*b.X, *b.Y)
}
