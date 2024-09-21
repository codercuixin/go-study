package main

import (
	"fmt"
	"sort"
)

func main() {
	// basic()
	// sortBySliceAndLessFunction()
	// sortByInterface()
	// binarySearch()
	reverse()
}


func reverse(){
	q := Queue{
		{text: "a", index: 5},
		{text: "b", index: 2},
		{text: "c", index: 6},
		{text: "d", index: 3},
		{text: "e", index: 1},
		{text: "f", index: 4},
	}
	sort.Sort(sort.Reverse(q))
	fmt.Println(q)
	
}

func binarySearch(){
	q := Queue{
		{text: "a", index: 5},
		{text: "b", index: 2},
		{text: "c", index: 6},
		{text: "d", index: 3},
		{text: "e", index: 1},
		{text: "f", index: 4},
	}

	sort.Sort(q)
	fmt.Println(q)

	i := sort.Search(len(q), func(i int)bool{
		return q[i].index > 6
	})
	fmt.Println("index> 6", i)

	i = sort.Search(len(q), func(i int)bool{
		return q[i].index >=3
	})
	fmt.Println("index>= 3", i)

}

func sortByInterface(){
	q := Queue{
		{text: "a", index: 5},
		{text: "b", index: 2},
		{text: "c", index: 6},
		{text: "d", index: 3},
		{text: "e", index: 1},
		{text: "f", index: 4},
	}
	fmt.Println(sort.IsSorted(q))

	sort.Sort(q)
	fmt.Println(sort.IsSorted(q))
	fmt.Println(q)
}

func basic() {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s)
}

func sortBySliceAndLessFunction(){
	s := [] struct{
		id int 
		name string 
	}{
		{5, "a"},
		{2, "b"},
		{6, "c"},
		{3, "d"},
		{1, "e"},
		{4, "f"},
	}

	sort.Slice(s, func(i, j int)bool{
		return s[i].id > s[j].id //倒序
	})

	fmt.Println(s)
}


type Data struct{
	text string 
	index int 
}

type Queue []Data 
func (q Queue)Len()int{
	return len(q)
}

func (q Queue)Less(i, j int) bool{
	return q[i].index < q[j].index
}

func (q Queue)Swap(i, j int){
	q[i], q[j] = q[j], q[i]
}