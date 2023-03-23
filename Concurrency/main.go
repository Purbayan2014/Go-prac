package Concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for indx := 1; indx <= 15; indx++ {
		fmt.Println("foo : ", indx)
	}

	// add the done to indicate that the subroutine foo is finished
	wg.Done()
}

func boo() {
	for indx := 1; indx <= 15; indx++ {
		fmt.Println("boo : ", indx)
	}
}

func do(x int) int {
	return x
}

func Concurrency() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// launching this function with its own go routine
	// 1 thing that we are waiting for that is that foo package
	wg.Add(1)
	go foo() // this foo function will run concurrently no need to wait for it
	boo()
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	ch := make(chan int)
	go func() {
		// adding the value into a channel
		ch <- do(2)
	}()
	// returning the value from the channel
	fmt.Println(<-ch)
	wg.Wait() // add this wait function to tell him to wait for
	// execution of the foo function

}
