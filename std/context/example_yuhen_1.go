package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	request(context.Background())
	fmt.Scanln()
}

func request(ctx context.Context){
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	resp := make(chan int)
	go handle(ctx, resp)

	//do something ....
	select {
	case v := <-resp:
		println(v)
	case <- ctx.Done():
		println(ctx.Err().Error())
	}
}

func handle(ctx context.Context, resp chan <- int){
	println("1/3 handle")
	cache(ctx, resp)
}

func cache(ctx context.Context, resp chan <-int){
	println("2/3 cache")

	//mock working
	time.Sleep(time.Second * 2)

	database(ctx, resp)
}

func database(ctx context.Context, resp chan <- int){
	//check done
	select{
	case <- ctx.Done():
		println("3/3 database: timeout")
		return
	default:
	}
	println("3/3 database")
	resp <- 100
}