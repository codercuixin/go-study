package main 

import "testing"
func assert(t *testing.T, b bool ){
	t.Helper()
	if !b {
		t.Fatal("assert fial ")
	}
}

func TestHelper(t *testing.T){
	assert(t, false)
}