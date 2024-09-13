package main

import (
	"fmt"
	"reflect"
	"unsafe"
)
func main(){
	// basic()
	// slice_len()
	// slice_len2()
	// slice_construct()
	// testIsInited()
	// dontsuppotCmp()
	// clearSlice()
	// bidirectionConvert()
	// pointer()
	// pointerOp()
	// useSliceAsFunctionArg()
	// jaddedArray()
	// testAppend()
	testCopy()
}
func testCopy(){
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	src := s[5:8]
	dst := s[4:] 
	//copy 使用同一个底层数组
	n := copy(dst, src) //s to [0, 1, 2, 3, 5, 6, 7, 7, 8, 9]
	fmt.Println(n, s)

	//在不同的数组间复制
	dst = make([]int, 6)
	n = copy(dst, src) 
	fmt.Println(n, dst) //dst is: [6, 7, 7, 0, 0, 0]

	b := make([]byte, 3)
	n = copy(b, "abcde")
	fmt.Println(n, b) // b is: [97 98 99]
}

func testAppend(){
	a := [...]int{0, 1, 2, 3, 99:0}
	fmt.Printf("a: %p ~ %p\n", &a[0], &a[len(a)-1])

	s := a[:4:8]
	pslice("s", &s)

	//小于等于 s.cap 限制
	s = append(s, []int{4, 5, 6}...)
	pslice("a", &s)

	// 大于 s.cap 限制
	s = append(s, []int{7, 8}...)
	pslice("a", &s)

	fmt.Println("a:", a[:len(s)])
	fmt.Println("s:", s)
	// 防止底层数组增长，一步到位。
	// s := make([]int, 0, 10000)

}
func pslice(name string, p*[]int){
	fmt.Printf("%s: %#v\n", name, *(*reflect.SliceHeader)(unsafe.Pointer(p)))
}


func jaddedArray(){
	s := [][]int{
		{1, 2},
		{10, 20, 30},
		{100},
	}
	s[1][2] +=100
	fmt.Println(s)
}


func useSliceAsFunctionArg(){
	s := []int{1,2,3}
	fmt.Println("addr: %p", &s[0])
	// only slice backed array adrress, len, cap three fields are copied 
	println(sum(s))
}

func sum(s[]int)(n int){
	fmt.Println("addr: %p", &s[0])
	for _, v:=range s{
		n += v
	}
	return
}

func pointerOp(){
	var a [4]int = [...]int{0, 1, 2, 3}
	var p *[4] int = &a
	//基于数组指针创建 slice 
	var s[]int = p[:]
	println(&s[2] == &a[2])

	//基于非数组指针创建切片
	// p2 := (*byte)(unsafe.Pointer)(&a[2])
	// var s2[]byte = unsafe.Slice(p2, 8)
	// fmt.Println(s2)
	
}

func pointer(){
	//不能以切片指针的方式访问元素，理由很简单，
	//因为 append 可能底层指向的数组已经是新的了, 也就是新地址了
	a :=[...]int{0, 1, 2, 3}
	s := a[:]

	slice_pointer := &s //切片指针
	element_pointer := &s[1] 

	//invalid operation: cannot index slice_pointer (variable of type *[]int)
	// _ = slice_pointer[1]

	_ = (*slice_pointer)[1]
	*element_pointer += 100

	fmt.Println(element_pointer == &a[1])
	fmt.Println(a)


	
}

func bidirectionConvert(){
	var a [4]int = [...]int{1, 2,3, 4}
	// array to slice, 指向原数组
	var s []int = a[:]
	println(&s[0] == &a[0]) // true
	
	//slice to array, 复制底层数组（片段）
	a2 := [4]int(s)
	println(&a2[0] == &a[0]) // false

	//slice to array:部分数据 1,2
	a3 := [2]int(s)
	println(a3[1] == a[1]) //true
	fmt.Println(a3)//false

	//slice to *array, 
	p2 := (*[4]int)(s)
	println(p2 == &a) //true
}

func clearSlice(){
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[:2:6]
	fmt.Printf("%v, len:%d, cap:%d, ptr:%p\n",
	 a, len(s), cap(s), &s[0])
	clear(s)
	fmt.Printf("%v, len:%d, cap:%d, ptr:%p\n",
	 a, len(s), cap(s), &s[0])
}

func dontsuppotCmp(){
	// s1 := []int{1,2}
	// s2 := []int{1,2,3}

	//invalid operation: s1 == s2 (slice can only be compared to nil)
	// println(s1 == s2)

}

func slice_construct(){
	//基于数据切片
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[:]
	fmt.Println(s, len(s), cap(s))
	// s1 :=a[2:5]
	// fmt.Println(s1, len(s1), cap(s1))
	// s2 :=a[2:5:7]
	// fmt.Println(s2, len(s2), cap(s2))
	// s3 :=a[2:5:7]
	// fmt.Println(s3, len(s3), cap(s3))
	// s4 :=a[4:]
	// fmt.Println(s4, len(s4), cap(s4))
	// s5 :=a[:4]
	// fmt.Println(s5, len(s5), cap(s5))
	// s6 :=a[:4:6]
	// fmt.Println(s6, len(s6), cap(s6))

	//按初始值创建slice，自动分配底层数组
	s1 := []int{1, 2, 4:4}
	fmt.Println(s1, len(s1), cap(s1))

	//自动创建底层数组。 指定 len，默认 cap 和 len 一致。
	s2 := make([]int, 5)
	fmt.Println(s2, len(s2), cap(s2))

	//自动创建底层数组。指定 len,cap 
	s3 := make([]int, 0, 5)
	fmt.Println(s3, len(s3), cap(s3))
}

func testIsInited(){
	var s1 []int ; //仅仅分配 header 内存，为初始化数据
	s2 := []int{}; // 初始化 slice
	s3 := make([]int, 0) //调用 make 初始化
	printIsInited(s1)
	printIsInited(s2)
	printIsInited(s3)
}

func printIsInited(s []int){
	fmt.Printf("%t, %d, %#v\n", s==nil, 
	unsafe.Sizeof(s), (*reflect.SliceHeader)(unsafe.Pointer(&s)))
}

func slice_len2(){
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:6:8]

	s1 := s[0:2:4]
	fmt.Println(s1, len(s1), cap(s1))

	s2 := s[1:6]
	fmt.Println(s2, len(s2), cap(s2))
}

func slice_len(){
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:6:8]
	//panic: runtime error: index out of range [5] with length 4
	// fmt.Println(s[5])

	for i, x := range s{
		fmt.Printf("s[%d]: %d\n", i, x)
	}
}

func basic(){
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:6:8]
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	fmt.Printf("a:%p ~ %p\n", &a[0], &a[len(a)-1])
	fmt.Printf("s:%p ~ %p\n", &s[0], &s[len(s)-1])
}