package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	b, _ := os.ReadFile("../../go.mod")
	h := sha256.New()
	h.Write(b)
	fmt.Printf("%x\n", h.Sum(nil))
}
