package ninja10

import (
	"fmt"
)

func Ninja10() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	cs := make(chan int)
	cs2 := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	go func() {
		cs2 <- 42
	}()

	fmt.Println(<-cs2)

	res := gen()
	receive(res)

	q := make(chan int)
	res2 := gen2(q)
	receive2(res2, q)

	newch := make(chan int)
	send100(newch)
	printch(newch)

	newch2 := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for indx := 0; indx < 10; indx++ {
				newch2 <- indx
			}
		}()
	}

	// print it
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-newch2)
	}
}

func send100(ch chan int) {
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()
}

func printch(ch chan int) {
	for val := range ch {
		fmt.Println("value from channel : ", val)
	}
}

func receive(c <-chan int) {
	for val := range c {
		fmt.Println(val)
	}
}

func receive2(a, q <-chan int) {
	for {
		select {
		case value := <-a:
			fmt.Println("value from channel a : ", value)
		case <-q:
			return
		}
	}
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func gen2(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for in := 0; in < 100; in++ {
			c <- in
		}
		q <- 1
		close(c)
	}()
	return c

}
