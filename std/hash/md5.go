package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	b, _ := os.ReadFile("../../go.mod")
	h := md5.New()
	h.Write(b)
	fmt.Printf("%x\n", h.Sum(nil))
}
