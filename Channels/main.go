package Channels

import (
	"fmt"
	"sync"
)

var signal sync.WaitGroup
var mutex sync.Mutex

func foo(c chan<- int) {
	// sending values through a range
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
	signal.Done()
}

func boo(c <-chan int) {
	fmt.Println(<-c)
	signal.Done()
}

func Channels() {
	// with no subroutine channels will block data
	// which creates a deadlock

	// channel that stores int
	c := make(chan int, 1) // 1 indicates the buffer
	go func() {
		mutex.Lock()
		c <- 22 // only one value added to buffer
		mutex.Unlock()
	}()
	fmt.Println(<-c)
	// two buffers
	ch := make(chan int, 2)
	go func() {
		mutex.Lock()
		ch <- 222
		ch <- 223
		mutex.Unlock()
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Unidirectional channels
	// send values only into the channels

	ch1 := make(chan<- int, 2)
	go func() {
		mutex.Lock()
		ch1 <- 2
		ch1 <- 22
		mutex.Unlock()
	}()

	// send out only
	ch2 := make(<-chan int, 2)
	fmt.Println(ch2)

	// using send and recv channels
	ch3 := make(chan int, 2)
	signal.Add(1)
	go foo(ch3)
	boo(ch3)

	signal.Wait()
	fmt.Println("All channels exited")

}
