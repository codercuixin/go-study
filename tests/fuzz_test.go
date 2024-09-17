package main

import (
	"testing"
)
//go test -v -fuzz Add -fuzztime 20s
func FuzzAdd(f *testing.F){
	f.Add(1, 1)
	f.Add(1, 2)
	f.Add(1, 3)
	f.Fuzz(func(t *testing.T, x int , y int ){
		addF(x, y)
	})
}

func addF(x, y int)int{
	return x + y
}