package main

import (
	"context"
	"time"
)

func request1(ctx context.Context){
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx = context.WithValue(ctx, "A", "a")
	go handle1(ctx)

	time.Sleep(1*time.Second)
}


func handle1(ctx context.Context){
	println("1/3 handle")
	ctx = context.WithValue(ctx, "B", "b")
	cache1(ctx)
}

func cache1(ctx context.Context){
	println("2/3 cache:")
	println("	", ctx.Value("A").(string))

	ctx = context.WithValue(ctx, "A", "a2")
	database1(ctx)
}

func database1(ctx context.Context){
	println("3/3 database")
	println("	", ctx.Value("A").(string))
	println("	", ctx.Value("B").(string))
}

func main(){
	//https://www.yuque.com/qyuhen/go/lpzddk
	request1(context.Background())
}
