package race_condition

import (
	"fmt"
	"runtime"
	"sync"
)

func RaceCondition() {
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("Goroutines : ", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(gs)
	for indx := 0; indx < gs; indx++ {
		go func() {
			// reading the counter
			// marks the critical section of the code so that
			// this part of the code is under lockdown
			// only the counter variable gets accessed in this part
			// and it doesnt forms a race condition
			mutex.Lock()
			v := counter
			// yielding it
			//time.Sleep(time.Second * 2) // sleeps for a bit after reading the value
			runtime.Gosched()
			// incrementing it
			v++
			counter = v
			mutex.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines : ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines : ", runtime.NumGoroutine())
	fmt.Println("Counter : ", counter)

}
