package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second * 2)
	defer t.Stop()

	for i:= 0;i<10; i++{
		fmt.Println(<-t.C)
	}
}