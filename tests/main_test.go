package main

import (
	"testing"
	"time"
)

func TestA(t *testing.T){
	time.Sleep(time.Second * 10)
}
func TestB(t *testing.T){
	time.Sleep(time.Second * 10)
}

func TestAdd(t *testing.T){
	z := add(1, 2)
	if z != 3{
		t.FailNow()
	}
}