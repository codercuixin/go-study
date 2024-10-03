package main

import (
	"fmt"
	"sync"
)
func gen(nums ...int) <- chan int{
	out := make(chan int)
	go func(){
		for _, n := range nums{
			out <- n 
		}
		close(out)
	}()
	return out 
}

func genWithDone(done <-chan struct{}, nums ...int) <- chan int{
	out := make(chan int)
	go func(){
		defer close(out)
		for _, n := range nums{
			select{
			case <-done:
				return
			case out <-n:
			}
			
		}
	}()
	return out 
}

func sq (in <- chan int) <- chan int{
	out := make(chan int)
	go func(){
		for n := range in {
			out <- n *n 
		}
		close(out)
	}()
	return out 
}

func sqWithDone ( done <-chan struct{}, in <- chan int) <- chan int{
	out := make(chan int)
	go func(){
		defer close(out)
		for n := range in {
			select{
			case out <- n *n :
			case <-done:
				return
			}
		}
	}()
	return out 
}

func merge(cs ...<-chan int) <-chan int{
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int){
		for n:= range c {
			out <-n 
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs{
		go output(c)
	}
	go func(){
		wg.Wait()
		close(out)
	}()
	return out 
}

func mergeWithDone(done <-chan struct{}, cs ...<-chan int) <-chan int{
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int){
		defer wg.Done()
		for n:= range c {
			select{
			case out <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs{
		go output(c)
	}
	go func(){
		wg.Wait()
		close(out)
	}()
	return out 
}

func main(){
	// c := gen(2, 3)
	// out := sq(c)
	// fmt.Println(<-out)
	// fmt.Println(<- out )

	// for n := range sq(sq(gen(2, 3))){
	// 	fmt.Println(n)
	// }

	// in := gen(2, 3)
	// c1 := sq(in)
	// c2 := sq(in)
	// for n:= range merge(c1, c2){
	// 	fmt.Println(n)
	// }

	done := make(chan struct{})
	defer close(done)
	in := genWithDone(done, 2,3)
	c1 := sqWithDone(done , in)
	c2 := sqWithDone(done , in)
	out := mergeWithDone(done, c1, c2)
	fmt.Println(<-out)
	//done will be closed by defered call

}
