package main 
import "fmt"
import "unsafe"

func printMap(m map[string]int){
	fmt.Printf("%t, %x\n", m==nil,
		*(*uintptr)(unsafe.Pointer(&m)))
}

func main(){
	//  basic()
	// testInit()
	// testExists()
	// testDelete()
	// testClear()
	// iteraterInRandom()
	// updateValueWithPointer()
	// updateValueWithWholeReplace()
	// editInIterate()
	concurrentReadWrite()
}

func concurrentReadWrite(){
	m := make(map[string]int)
	
	//wrtie
	go func(){
		for{
			m["a"] +=1
		}
	}()

	//read
	go func(){
		for{
			_ = m["b"]
		}
	}()

	select{}
}

func editInIterate(){
	m := make(map[int]int)
	for i:=0; i<10; i++{
		m[i] = i+10
	}
	for k:= range m{
		if k==5{
			m[100] = 1000
		}
		delete(m, k)
		fmt.Println(k, m)
	}

}

func updateValueWithWholeReplace(){
	type user struct{
		name string
		age byte 
	}

	m := map[int]user{
		1: user{"u1", 19},
	}
	m[1] =user{
		"u1",
		20,
	}
	fmt.Println(m[1])
}

func updateValueWithPointer(){
	type user struct{
		name string
		age byte 
	}

	m := map[int]*user{
		1: &user{"u1", 19},
	}
	m[1].age = 20
	fmt.Println(m[1])
}

func iteraterInRandom(){
	m := make(map[int]struct{})
	for i:=0; i<10; i++{
		m[i] = struct{}{} 
	}
	for i:=0; i<4; i++{
		for k, _ := range m{
			print(k, ",")
		}
		println()
	}
}

func testClear(){
	m := map[string]int{
		"a":1,
		"b":2,
	}
	println(len(m))
	clear(m)
	println(len(m))
}

func testDelete(){
	var m map[string]int //nil
	_ = m["a"]
	delete(m, "a")
	//panic: assignment to entry in nil map
	// m["a"] = 0
}

func testExists(){
	m := map[string]int{}

	x := m["a"]
	fmt.Println(x) //0

	x, ok := m["a"]
	fmt.Println(x, ok)//0, false

	m["a"] = 0
	x, ok = m["a"]
	fmt.Println(x, ok)//0, true
}

func mapCannotCmp(){
	// m1 := map[int]int{}
	// m2 := map[int]int{}
	
	// invalid operation: m1 == m2 (map can only be compared to nil)compilerUndefinedOp
	// _ = m1 == m2 
}

func testInit(){
	m1 := map[string]int{
		"a":1,
		"b":2,
	}
	//省略符合类型标签
	m2 := map[string]struct{
		id int
		name string 
	}{
		"a":{1, "u1"},
		"b":{2, "u2"},
	}

	//分配足够初始容量，比量后续扩张
	m3 := make(map[string]int, 10)
	m3["a"] = 1

	fmt.Println(m1, len(m1))
	fmt.Println(m2, len(m2))
	fmt.Println(m3, len(m3))

}

func basic(){
	var m1 map[string]int
	m2 := map[string]int{}
	m3 := make(map[string]int, 0)
	printMap(m1)
	printMap(m2)
	printMap(m3)
}