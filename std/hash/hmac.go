package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	key := []byte("password")
	h := hmac.New(sha256.New, key)
	h.Write([]byte("hello world"))
	fmt.Printf("%x\n", h.Sum(nil))
}
