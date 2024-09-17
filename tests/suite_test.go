package main

import "time"
import "testing"

func TestA2(t *testing.T) {
	time.Sleep(time.Second)
}

func TestB2(t *testing.T){
	time.Sleep(time.Second)
}

func TestC2(t *testing.T){
	time.Sleep(time.Second)
}

//go test -v -run "Suite"
//go test -v -run "Suite/A"
func TestSuite(t *testing.T){
	t.Log("setup")
	defer t.Log("teardown")
	t.Run("A", TestA2)
	t.Run("B", TestB2)
	t.Run("C", TestC2)
}