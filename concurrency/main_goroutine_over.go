package main

import "time"
func main(){
	go func(){
		defer println("g done")
		time.Sleep(time.Second)
	}()

	defer println("main goroutine done")
}