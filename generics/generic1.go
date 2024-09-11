package main

import "fmt"

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
	~string
}

func maxGenerics[T ordered](s1 []T) T {
	if len(s1) == 0 {
		panic("slice is empty")
	}

	max := s1[0]
	for _, v := range s1[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

type myString string

func main() {
	var m int = maxGenerics([]int{1, 2, -4, -6, 7, 0})
	fmt.Println(m)
	fmt.Println(maxGenerics([]string{"11", "22", "44", "66", "77", "10"}))  // 输出：77
	fmt.Println(maxGenerics([]float64{1.01, 2.02, 3.03, 5.05, 7.07, 0.01})) // 输出：7.07
	fmt.Println(maxGenerics([]int8{1, 2, -4, -6, 7, 0}))                    // 输出：7
	fmt.Println(maxGenerics([]myString{"11", "22", "44", "66", "77", "10"}))
	foo1[int,](5, "hello")

	var a int = foo2[int](5)
	print(a)


	var s1 = maxableSlice[int]{
		elems:[]int{1, 2, -4, -6},
	}
	print(s1)

	type maxAlias = maxableSlice[int]
	var s2 = maxAlias{
		elems: []int{1,2, 3},
	}
	print(s2)
}


func foo1[T comparable, E any](a int, s E){

}

func foo2[T any](a int)T{
	var zero T 
	return zero 
}

type maxableSlice[T ordered] struct{
	elems  []T 
}

func(s1 *maxableSlice[T]) max() T{
	if len(s1.elems) == 0{
		panic("slice is empty")
	}
	max := s1.elems[0]
	for _, v := range s1.elems[1:]{
		if v > max{
			max = v 
		}
	}
	return max 
}

type Set[T comparable] map[T]struct{}
