package main

import (
	"errors"
	"log"
)

func do(x int) error {
	if x <= 0 {
		return errors.New("x <= 0")
	}
	return nil
}

func main() {
	// if1()
	// if_fail_fast()
	// switch1()
	// switch_implict_break()
	// switchConst()
	// switchVar()
	// fallThroughTest()
	// fallThroughTest2()
	// fallThroughTest3()
}

func switchIgnoreConditionalVar(){
	switch x:=1; {
	case x == -1, x==1:
		println("first")
	case x > 1 && x<=10:
		println("second")
	case x> 10:
		println("b")
	default:
		println("d")
	}
}

func fallThroughTest() {
	switch x := 5; x {
	case 5:
		println("5")
		fallthrough
	case 6:
		println("6")
	case 7:
		println("7")
	}
}


func fallThroughTest2() {
	switch x := 5; x {
	default:
		println("default")
	case 5:
		println("5")
		//cannot fallthrough final case in switch
		// fallthrough
	}
}

func fallThroughTest3() {
	switch x := 5; x {
	case 5:
		x += 10
		if x < 15 {
			break
		}
		fallthrough
	default:
		println("default")
	}
}

func switchConst() {

	x := 2
	switch x {
	case 1:
		println("1")
	case 2:
		println("2")
	}
}

func switchVar() {
	a, b := 1, 2
	x := a
	switch x {
	case b, a:
		println("b, a")
	case a:
		println("a")
	}
}

func switch_implict_break() {
	switch x := 1; x {
	case 0: // 隐式 break 中断
	case 1:
	case 2:
		println("b")
	}
}

func switch1() {
	a, b, c := 1, 2, 3
	switch x := 5; x {
	case a, b:
		println("a or b")
	case c:
		println("c")
	case 4:
		println("d")
	default:
		println(x)
	}
}

func if1() {
	x := 10
	// 局部变量 err 作用域覆盖所有分支
	if err := do(x); err == nil {
		x++
		println(x)
	} else {
		log.Fatalln(err)
	}
}

func if_fail_fast() {
	x := 10
	err := do(x)
	// 局部变量 err 作用域覆盖所有分支
	if err != nil {
		log.Fatalln(err)
	}
	x++
	println(x)
}
