package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t := time.NewTimer(time.Second)
	fmt.Println(<-t.C)

	c := time.After(time.Second)
	fmt.Println(<-c)

	var wg sync.WaitGroup
	wg.Add(1)
	t = time.AfterFunc(time.Second*3, func(){
		defer wg.Done()
		println("after func")
	})
	t.Stop()
	// wg.Wait()
	println("done")
}