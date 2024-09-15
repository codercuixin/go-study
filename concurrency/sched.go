package main

import (
	"runtime"
	"sync"
	"time"
)

func main(){
	// hangup()
	// siganlRunners()
	executeInOrders()
}

func executeInOrders(){
	const COUNT = 5 
	var wg sync.WaitGroup
	wg.Add(COUNT)

	var chans [COUNT]chan struct{}
	for i:=0; i<COUNT; i++{
		chans[i] = make(chan struct{})

		go func(id int){
			defer wg.Done()
			<- chans[id]
			println("done", id)
		}(i)
	}
	for _, x := range []int{4, 0, 1, 3, 2}{
		close(chans[x])
		time.Sleep(time.Millisecond * 10)
	}
	wg.Wait()
}

func siganlRunners(){
	var wg sync.WaitGroup
	signal := make(chan struct{})
	for i := 0; i<10; i++{
		wg.Add(1)

		go func(id int){
			defer wg.Done()
			<-signal //wait for siganl
			println("runner ", id, " done.")
		}(i)
	}
	close(signal)
	wg.Wait()
}

func hangup(){
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan struct{}), make(chan struct{})

	go func(){
		defer wg.Done()
		<- a 
		for i:=0; i<5; i++{
			println("a", i)
		}
	}()
	go func(){
		defer wg.Done()
		<- b 
		for i:=0; i<5; i++{
			println("b", i)
			if i == 2{
				//让其他 go routine 运行下。
				//Gosched yields the processor, allowing other goroutines to run. It does not
// suspend the current goroutine, so execution resumes automatically.
				runtime.Gosched()
			}
		}
	}()
	//make b run first 
	close(b)
	close(a)
	wg.Wait()
}