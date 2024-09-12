package main 
import "fmt"
type TestError struct{
	x int 
}

func(e *TestError) Error()string{
	return fmt.Sprintf("test:%d", e.x)
}

var ErrZero = &TestError{x: 0}

func main(){
	// var e error = ErrZero
	// fmt.Println(e == ErrZero)

	// if t, ok := e.(*TestError); ok{
	// 	fmt.Println(t.x)
	// }

	err := nilVsError()
	println(err == nil)
}

func nilVsError() error{
	var err *TestError
	println(err == nil)
	// 类型转换为 error 接口, 其实是带类型为 TestError 的
	// 只有类型和值都为 nil 时，才等于 nil。
	return err 
}