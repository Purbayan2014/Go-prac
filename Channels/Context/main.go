package Context

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func Context2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancelling only when finished

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			// defer cancel hits this and cancels the job
			case <-ctx.Done():
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()

	return dst
}

func Context() {
	// context created with context.Background()
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("error check 1: ", ctx.Err())
	fmt.Println("num goroutines 1: ", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			// done can be used in select statements
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num goroutines 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")
	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num goroutines 3:", runtime.NumGoroutine())

}
