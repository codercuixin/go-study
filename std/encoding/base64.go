package main

import (
	"bytes"
	"encoding/base64"
)

func main() {
	b := bytes.NewBuffer(nil)
	w := base64.NewEncoder(base64.StdEncoding, b)
	w.Write([]byte("1234"))
	w.Write([]byte("abc"))
	w.Close()
	// !!!!
	println(b.String())
	r := base64.NewDecoder(base64.StdEncoding, b)
	bs := make([]byte, base64.StdEncoding.DecodedLen(b.Len()))
	r.Read(bs)
	println(string(bs))
}
