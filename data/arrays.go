package main 
// import "unsafe"
import "fmt"
func main(){
	// var d1 [unsafe.Sizeof(0)]int // 
	// var d2 [2]int // 元素类型相同，长度不同
	// //cannot use d2 (variable of type [2]int) as [8]int value in assignment
	// d1 = d2 

	//  good_init()
	// init_struct()
	// muldim_array()
	// equal()
	// pointer()
	// pointerArray()

	testCopy()
}

func testCopy(){
	d := [...]byte{1, 2, 3}
	d2 := d // copy array
	d3 := *(&d) // copy array
	fmt.Printf(" d: %p\n", &d)
	fmt.Printf("d2: %p\n", &d2)
	fmt.Printf("d3: %p\n", &d3)

	d4 :=valCopyArray(d)
	fmt.Printf("valCopyArray. return: %p\n", &d4)

	p := ptrCopyArray(&d)
	fmt.Printf("ptrCopyArray. return: %p\n", p)

	fmt.Printf("d: %v\n", d)

}



func valCopyArray(d [3]byte) [3]byte{
	fmt.Printf("val: %p\n", &d)
	d[0] += 100
	return d 
}

func ptrCopyArray(p *[3]byte) *[3]byte{
	fmt.Printf("ptr: %p\n", p)
	p[0] += 100
	return p
}




func pointer(){
	d := [...]int{0, 1, 2, 3}
	var p *[4]int = &d  //数组指针
	var pe *int = &d[1] //元素指针

	p[0] += 10
	*pe += 20
	fmt.Println(d)
}

func pointerArray(){
	a, b  := 1, 2 
	d := [...]*int{&a, &b}
	*d[1] += 10
	fmt.Println(d)
	fmt.Println(a, b)
}





// 相等比较，需元素支持
func equal(){
	var a, b [2]int
	println(a == b)

	c := [2]int{1, 2}
	d := [2]int{0, 1}
	println(c == d)

}

func muldim_array(){
	x := [...][2]int{
		{1, 2},
		{3, 4},
	}
	y := [...][3][2]int{
		{
			{1, 2},
			{3, 4},
			{5, 6},
		},
		{
			{11, 12},
			{13, 14},
			{15, 16},
		},
	}
	fmt.Println(x, len(x), cap(x)) //2, 2
	fmt.Println(y, len(y), cap(y)) //2 , 2
}

func init_struct(){
	type user struct{
		id int 
		name string 
	}
	var users  = [...]user{
		{1, "zs"},
		{2, "ls"},
	}
	fmt.Println(users)
}

func good_init(){
	var a [4]int //元素默认初始化为 0
	b := [4]int{2, 5} //未提供的元素，默认为 0
	c := [4]int{5, 3:10} //可指定索引位置初始化
	d := [...]int{1,2,3} // 编译器按初始化之数量确定数组长度
	e := [...]int{10, 3:100} //支持索引初始化，数组长度与此相关
	fmt.Println(a, b, c, d, e)
}
