package main

import "fmt"

func main() {
	// fmt.Printf("% X\n", "abc")
	// fmt.Printf("%02X\n", []int{0x1, 0xF2})
	// fmt.Printf("%02X\n", []int{0x1222, 0xF2})
	// fmt.Printf("%-5s\n", "abc")
	// fmt.Printf("%010.4f", 12.3456789)

	// fmt.Printf("%[1]d %d %[1]d %[3]d", 1, 2, 3)
	// var s = struct{
	// 	name string
	// 	age int
	// }{
	// 	name: "hello world",
	// 	age : 18,
	// }
	// fmt.Printf("%T", s)
	//default format
	// fmt.Printf("%v, %d", s, 10)
	// 输出完整的结构体信息，带字段名
	// fmt.Printf("%+v", s)
	
	//按 go 语法格式输出
	// fmt.Printf("%#v", s)

	//指针
	// fmt.Printf("%p", &s)
	// fmt.Printf("%b\n", 10)
	// fmt.Printf("%d\n", 10)
	// fmt.Printf("%x\n", 10)//a
	// fmt.Printf("%X\n", 10) //A

	// fmt.Printf("%f\n", 1.2) //1.200000
	// fmt.Printf("%s", "hello")
	// + 添加数字符号
	// fmt.Printf("%+d", -100)
	// - 左对齐
	// fmt.Printf("%-3s", "a")
	//增加前缀
	// fmt.Printf("%#b", 10) //0b1010
	// fmt.Printf("%#x", 10) //oxa
	// fmt.Printf("%#o", 10) //012

	// fmt.Printf("%04d", 10); //填充 0， 0010
	// fmt.Printf("% 4d", 10); //填充空格，“  10”

	// err := fmt.Errorf("%d %s", 10, "hello world")
	// fmt.Println(err)

	// fmt.Printf("%t", true)
	// fmt.Printf("%c", '你')
	// fmt.Printf("%o", 10)//base 8
	// fmt.Printf("%O", 10)//base 8 with 0o prefix
	// s := `He said, "Hello, world!"\n`
 	// formatted :=	fmt.Sprintf("%q", s)
	// fmt.Println(formatted) //"He said, \"Hello, world!\"\\n"

	// fmt.Printf("%U\n", 0x78)
	// fmt.Printf("%04U\n", 0x78)
	// fmt.Printf("%U\n", '你')
	// fmt.Printf("%b\n",  3.1415926)
	// fmt.Printf("%b\n",   100.0)
	// fmt.Printf("%e\n", 100.0)

	// fmt.Printf("%f", 123.456)
	fmt.Printf("%g\n", 1334777777777.0)
	fmt.Printf("%g\n", 122.0)
}
