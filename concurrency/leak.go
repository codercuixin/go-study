package main

import (
	"runtime"
	"time"
)

func test() chan byte {
	c := make(chan byte)
	go func() {
		buf := make([]byte, 0, 10<<20)
		for {
			d, ok := <-c
			if !ok {
				return
			}
			buf = append(buf, d)
		}
	}()
	return c
}

func main() {
	//GODEBUG=gctrace=1 go run leak.go 
	//GODEBUG="schedtrace=10000,scheddetail=1" go run leak.go 
	for i := 0; i < 5; i++ {
		test()
	}
	for {
		time.Sleep(time.Second)
		runtime.GC()
	}
}