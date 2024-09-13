package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main(){
	//  basic()
	// emptyString()
	// rawString()
	// strcomp()
	// substring()
	// iterateString()
	// useInInnerMethod()
	// convert()

	// contact()
}

func contact(){
	 s := "新年"
	 s2 := string(s[0:2] + s[4:])
	 println(s2)
	fmt.Println(utf8.ValidString(s2))
}


func convert(){
	var r rune = '我' //Unicode 字符
	var s string = string(r)
	var b byte = byte(r)
	var s2 string = string(b)
	var r2 rune = rune(b)
	fmt.Printf("%c, %U\n", r, r)
	fmt.Printf("%s, %X, %X, %X\n", s, b, s2, r2)
}

func useInInnerMethod(){
	s := "de"
	bs := make([]byte, 0)
	bs = append(bs, "abc"...)
	bs = append(bs, s...)

	buf := make([]byte, 5)
	copy(buf, "abc")
	copy(buf[3:], s)


	fmt.Printf("%s\n", bs)
	fmt.Printf("%s\n", buf)
}

func iterateString(){
	s := "新年好"
	//byte iterate
	for i:=0; i<len(s); i++{
		fmt.Printf("%d: %X\n", i, s[i])
	}

	//rune/unicode code point
	for i, v := range s{
		fmt.Printf("%d: %U\n", i, v)
	}
}

func substring(){
	//just beacuse all in ascill, one byte to one unicode code point
	s := "hello world!"
	s2 := s[:4]
	p1 := (*reflect.StringHeader)(unsafe.Pointer(&s))
	p2 := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Printf("%#v, %#v\n", p1, p2)
}


func strcomp(){
	s := "ab"+
	"cd"
	println(s == "abcd")
	println(s > "abc")
}

func rawString(){
	s := ` this is a line\r\n
	line2
	hello
	world
	`
	println(s)
}

func emptyString(){
	var s string 
	println( s == "")
}

func basic(){
	s := "新年好\x61\142\u0041"

	bs := []byte(s)
	rs := []rune(s) // rune/int32 unicode code point 

	fmt.Printf("% X, %d\n", s, len(s)) // return byte length 12
	fmt.Printf("% X, %d\n", bs, utf8.RuneCount(bs)) // return unicode length 6
	fmt.Printf("%U, %d\n", rs, utf8.RuneCountInString(s))

	fmt.Printf("%X\n", s[1]) // s[1] return byte 1
}