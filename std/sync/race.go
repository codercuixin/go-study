package main

import "sync"

// build with race detection,  then run
// go build -race && ./sync
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	x := 0
	go func(){
		defer wg.Done()
		x++
	}()

	go func(){
		defer wg.Done()
		println(x)
	}()
	wg.Wait()
}