package FanOut

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func FanOut() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go FanOutIn(c1, c2)

	for val := range c2 {
		fmt.Println(val)
	}
	fmt.Println("about to exit")
}

func FanOutIn(c1 chan int, c2 chan int) {
	var wg sync.WaitGroup
	for va := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(va)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(v2 int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(5000)))
	return v2 + rand.Intn(1000)
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
