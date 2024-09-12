package main

import (
	"log"
	"runtime"
)

func main() {
	// recoverFromPanic()
	// recoverLastPanic()
	// recoverTwoPanic()
	// recoverRight()
	nopanic()
	panicnil()
}

func panicnil(){
	defer func(){
		e := recover()
		println(e == nil)
		_, ok := e.(*runtime.PanicNilError)
		println(ok)
	}()
	panic(nil)
}

func nopanic(){
	defer func(){
		e := recover()
		println(e == nil)
	}()
}

func recoverRight(){
	// defer catch() //有效，在延迟函数内调用
	// defer log.Println(recover()) // 无效，作为参数立刻执行
	defer recover() //无效，被直接注册为延迟调用
	panic("panic")
}
func catch(){
	recover()
}

func recoverTwoPanic(){
	defer func() {
		log.Println(recover()) //捕获第二次 panic
	}()
	func(){
		defer func(){
			panic("panic2") //第二次 panic
		}()
		defer func(){
			log.Println(recover()) // 捕获第一次 panic	
		}()
		panic("panic1") //第一次 panic
	}()
	println("exit")
}

func recoverLastPanic(){
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	func(){
		defer func(){
			panic("panic2")
		}()
		panic("panic1")
	}()
	println("exit")
}

func recoverFromPanic(){
	defer func() {
		// 拦截恐慌，返回数据
		// 数据未必是 error，也可能是 nil
		// 无法返回 panic 的后续位置继续执行
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	func(){
		panic("p1")
	}()

	println("after panic, exit")
}