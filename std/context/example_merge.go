package main

import (
	"context"
	"errors"
	"fmt"

)

func main() {
	mergeCancel := func(ctx, cancelCtx context.Context)(context.Context, context.CancelFunc){
		ctx, cancel := context.WithCancelCause(ctx)
		stop := context.AfterFunc(cancelCtx, func() {
			cancel(context.Cause(cancelCtx))
		})
		return ctx, func() {
			stop()
			cancel(context.Canceled)
		}
	}

	ctx1, cancel1 := context.WithCancelCause(context.Background())
	// defer cancel1(errors.New("ctx1 canncelled"))

	ctx2, cancel2 := context.WithCancelCause(context.Background())

	mergedCtx, mergedCancel := mergeCancel(ctx1, ctx2)
	defer mergedCancel()
	cancel1(errors.New("ctx1 canncelled"))
	defer cancel2(errors.New("ctx1 canceled"))
	<- mergedCtx.Done()
	fmt.Println(context.Cause(mergedCtx))



}