package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	var rw sync.RWMutex
	x := 0

	wg.Add(1)
	go func(){
		defer wg.Done()
		for i:=0; i<5; i++{
			rw.Lock()
			time.Sleep(time.Second)
			now := time.Now().Format("15:04:05")
			x++
			fmt.Printf("[W] %d,  %v\n", x, now)
			rw.Unlock()
		}
	}()

	for i:=0; i<5; i++{
		wg.Add(1)
		go func(id int){
			defer wg.Done()
			rw.RLock()
			time.Sleep(time.Second)
			now := time.Now().Format("15:04:05")
			fmt.Printf("[R] %d,  %v\n", x, now)
			rw.RUnlock()
		}(i)
	}
	wg.Wait()
}