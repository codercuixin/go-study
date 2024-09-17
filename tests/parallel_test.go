package main

import (
	"testing"
	"time"
)

func TestA1(t *testing.T){
	t.Parallel()
	time.Sleep(time.Second * 10)
}

func TestB1(t *testing.T){
	t.Parallel()
	time.Sleep(time.Second * 10)
}

func TestC1(t *testing.T){
	time.Sleep(time.Second * 10)
}