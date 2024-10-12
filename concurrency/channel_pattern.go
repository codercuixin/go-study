package main

import (
	"sync"
	"time"
	"fmt"
)
func main(){
	// testNewRecv()
	// timeout()
	testSema()
}
type Pool[T any] chan T 

func NewPool[T any](cap int) Pool[T]{
	return make(chan T, cap)
}
func(p Pool[T]) Get()(v T, ok bool){
	select{
	case v = <-p: 
		ok = true
	default:
	}
	return
}


type Sema struct{
	c chan struct{}
}
func NewSema(n int) * Sema{
	return &Sema{
		c: make(chan struct{}, n),
	}
}

func(m *Sema) Accquire(){
	m.c <- struct{}{}
}

func(m *Sema) Release(){
	<-m.c 
}

func testSema(){
	var wg sync.WaitGroup
	sem := NewSema(2)
	for i:=0; i<5; i++{
		wg.Add(1)
		go func(id int){
			defer wg.Done()
			sem.Accquire()
			defer sem.Release()
			for n:=0; n <3; n++{
				time.Sleep(time.Second *2)
				fmt.Println(id, time.Now())
			}
		}(i)
	}

	wg.Wait()
}

func timeout(){
	quit := make(chan struct{})
	c := make(chan int)
	go func(){
		defer close(quit)
		select{
		case x, ok := <-c: 
			println(x, ok)
		case <- time.After(time.Second):
			return
		}
	}()
	<-quit 
}

func newRecv[T any](cap int)(data chan T, done chan struct{}){
	data = make(chan T,cap)
	done = make(chan struct{})
	go func(){
		defer close(done)
		for v:= range data{
			println(v)
		}
	}()
	return
}

func testNewRecv(){
	data, done := newRecv[int](3)
	for i:=0; i<10;i++{
		data <- i 
	}
	close(data)
	<- done 
}



