package main

import "fmt"

func a(){

}
func b(){

}


type FormatFunc func(string, ...any) (string)

func toString(f FormatFunc, s string, a ...any) string{
	return f(s, a...)
}




func main(){
	// 不支持命名函数嵌套
	// func add(x, y int)int {
	// 	return x+ y
	// }

	// println(a == nil)
	// println(b == nil)

	// var f func() = func(){
	// 	println("hello")
	// }

	// exec(f)
	println(toString(func(s string, a ...any) string {
		return fmt.Sprintf(s, a...)
	}, "hello %s %d", "world", 100))
}

func exec(f func()){
	f()
}

