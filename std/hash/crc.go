package main

import (
	"hash/crc32"
	"fmt"
	"os"
)

func main() {
	b, _ := os.ReadFile("../../go.mod")
	h := crc32.NewIEEE()
	h.Write(b)
	fmt.Printf("%x\n", h.Sum(nil))
}
