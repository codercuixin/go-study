package main

import (
	"context"
	"fmt"
	"time"
)
func main(){
	gen := func(ctx context.Context) <-chan int{
		dst := make(chan int)
		n :=1
		go func(){
			for{
				select{
				case <-ctx.Done():
					{
						fmt.Println("context has done")
						// 不 sleep 打印不出来
						time.Sleep(1*time.Microsecond)
						return
					}
				case dst <- n:
					n ++
				}
				
			}
		}()
		return dst 
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n:= range gen(ctx){
		fmt.Println(n)
		if( n== 5){
			break
		}
	}

}