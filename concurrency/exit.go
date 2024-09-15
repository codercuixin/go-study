package main

import (
	"os"
	"runtime"
	"time"
)

func main() {
	// callstackExit()
	// goExit()
	osExit()
}

func osExit(){
	go func(){
		defer println("g done")
		time.Sleep(time.Second)
	}()
	defer println("main done")
	// 不等待其他任务，也不执行延迟调用
	os.Exit(-1)
}



func goExit(){
	q := make(chan struct{})
	go func(){
		defer close(q)
		defer println("done.")
		time.Sleep(time.Second)
	}()
	//wait above go routine done 
	runtime.Goexit()
	<-q 
}




func callstackExit(){
	q := make(chan struct{})
	go func(){
		defer close(q)
		defer println("defer done.")
		a()
		
		b()
		// c will not be executed
		c()
	}()
	<-q 
}

func a() {
	println("a")
}

func b() {
	println("b")
	runtime.Goexit()
}

func c(){
	println("c")
}