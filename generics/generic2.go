package main

import (
	"fmt"
	"sync"
)

type Lockable[T any] struct {
	t T
	Slice[int]
	sync.Mutex
}

func (l *Lockable[T]) Get() T {
	l.Lock()
	defer l.Unlock()
	return l.t
}

func (l *Lockable[T]) Set(v T) {
	l.Lock()
	defer l.Unlock()
	l.t = v
}

type Slice[T any] []T

func (s Slice[T]) String() string {
	if len(s) == 0 {
		return ""
	}
	var result = fmt.Sprintf("%v", s[0])
	for _, v := range s[1:] {
		result = fmt.Sprintf("%v, %v", result, v)
	}
	return result
}

type Foo struct{
	Slice[int]
}


func main() {
	n :=Lockable[string]{
		t : "hello",
		Slice :[]int{1, 2, 3},
	}
	println(n.String())
	
	f := Foo{
		Slice: []int{1, 2, 3},
	}
	println(f.String())

}