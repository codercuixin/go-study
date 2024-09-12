package main

import (
	"errors"
	"fmt"
)

func database() error{
	return errors.New("data")
}

func cache() error{
	if err := database(); err !=nil{
		return fmt.Errorf("cache miss: %w", err)
	}
	return nil 
}

func handle() error{
	return cache()
}

func main(){
	// fmt.Println(handle())
	//打包
	a := errors.New("a")
	b := fmt.Errorf("b, %w", a)
	c := fmt.Errorf("c, %w", b)

	fmt.Println(c) // c, b, a
	fmt.Println(errors.Unwrap(c) == b) //true

	//递归检查
	fmt.Println(errors.Is(c, a)) //true

	x := &TestError2{}
	y := fmt.Errorf("y, %w", x)
	z := fmt.Errorf("z, %w", y)

	var x2 *TestError2 
	// 提取（二级指针）类型匹配的错误对象。
	if errors.As(z, &x2){
		fmt.Println(x == x2)
	}
}

type TestError2 struct{

}
func (*TestError2) Error()string{
	return ""
}