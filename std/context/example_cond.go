package main

import (
	"context"
	"sync"
	"time"
	"fmt"
)
func main(){
	waitOnCond := func(ctx context.Context, cond *sync.Cond, 
	conditionMet func()bool) error{
		stopf := context.AfterFunc(ctx, func(){
			cond.L.Lock()
			defer cond.L.Unlock()

			cond.Broadcast()
		})
		defer stopf()

		for !conditionMet(){
			cond.Wait()
			if ctx.Err() != nil{
				return ctx.Err()
			}
		}
		return nil 
	}

	cond := sync.NewCond(new(sync.Mutex))
	var wg sync.WaitGroup
	for i := 0; i<4; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Microsecond)
			defer cancel()
			cond.L.Lock()
			defer cond.L.Unlock()

			err := waitOnCond(ctx, cond, func()bool{
				return false 
			})
			fmt.Println(err)
		}()
	}
	wg.Wait()
}