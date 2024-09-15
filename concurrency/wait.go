package main

import (
	"context"
	"sync"
)

func main() {
	// emptyStructChannelWait()
	// waitGroupWait()
	// contextWait()
	lockWait()
}

func lockWait(){
	var lock sync.Mutex
	lock.Lock()
	go func(){
		defer lock.Unlock()
		println("sub go rountine done.")
	}()
	lock.Lock()
	lock.Unlock()
	println("parent done")
}

func contextWait(){
	ctx, cancel := context.WithCancel(context.Background())
	go func(){
		defer cancel()
		println("done.")
	}()
	<- ctx.Done()
}


func waitGroupWait() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i :=0; i<10; i++{
		go func (id int){
			defer wg.Done()
			println(id, "done")
		}(i)
	}
	// wait until all done
	wg.Wait()
	println("all done")
}

func emptyStructChannelWait() {

	q := make(chan struct{})

	go func() {
		defer close(q)
		println("done.")
	}()
	<-q
}
