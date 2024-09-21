package main

import (
	"fmt"
	"sort"
)
type Change struct{
	user string 
	language string
	lines int
}

type lessFunc func(p1, p2 *Change)bool 

type mutiSorter struct{
	changes []Change
	less []lessFunc 
}
func (ms *mutiSorter) Len() int{
	return len(ms.changes)
}
func (ms *mutiSorter) Swap(i, j int){
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}


func (ms *mutiSorter) Less(i, j int) bool{
	p, q := &ms.changes[i], &ms.changes[j]
	var k int 
	for k = 0; k < len(ms.less) - 1; k++{
		less := ms.less[k]
		if less(p,q){
			return true
		}else if less(q, p){
			return false 
		}
	}
	return ms.less[k](p,q)
}
func OrderBy(less...lessFunc) *mutiSorter{
	return &mutiSorter{
		less: less,
	}
}

func (ms *mutiSorter) Sort(changes []Change){
	ms.changes = changes 
	sort.Sort(ms)
}

func main(){
	var changes = []Change{
		{"gri", "Go", 100},
		{"ken", "C", 150},
		{"glenda", "Go", 200},
		{"rsc", "Go", 200},
		{"r", "Go", 100},
		{"ken", "Go", 200},
		{"dmr", "C", 100},
		{"r", "C", 150},
		{"gri", "Smalltalk", 80},
	}

	user := func(c1,c2 *Change) bool{
		return c1.user <c2.user 
	}
	language :=func(c1, c2*Change) bool{
		return c1.language <c2.language
	}
	OrderBy(user, language).Sort(changes);
	fmt.Println(changes)
}