package main 

import (
	"fmt"
	"time"
)
//修改签名
func partial(f func(int), x int) func(){
	return func() {
		f(x)
	}
}

//增加功能
func proxy(f func()) func(){
	return func() {
		n := time.Now()
		defer func(){
			fmt.Println(time.Now().Sub(n))
		}()
		f()
	}
}

func main(){
	test := func(x int){
		println(x)
	}
	var f func() = partial(test, 10)
	f()

	proxy(f)()
}
