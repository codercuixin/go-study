package main

import "sync"

func main() {
	var once sync.Once
	f1 := func(){
		println("1")
	}
	f2 := func(){
		println("2")
	}
	once.Do(f1)
	once.Do(f1)
	once.Do(f2)
	once.Do(f2)
}