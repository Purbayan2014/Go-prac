package Atomic

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func RaceCondition() {
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("Goroutines : ", runtime.NumGoroutine())
	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for indx := 0; indx < gs; indx++ {
		go func() {
			// reading the counter
			// marks the critical section of the code so that
			// this part of the code is under lockdown
			// only the counter variable gets accessed in this part
			// and it doesnt forms a race condition
			atomic.AddInt64(&counter, 1) // adding 1 to the counter
			runtime.Gosched()
			fmt.Println("Counter :", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines : ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines : ", runtime.NumGoroutine())
	fmt.Println("Counter : ", counter)

}
