package send_recv

import (
	"fmt"
	"sync"
)

func recv(e, o <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	//for {
	//	select {
	//	case v := <-e:
	//		fmt.Println("from the even", v)
	//	case v := <-o:
	//		fmt.Println("from the odd", v)
	//	case v, ok := <-q:
	//		if !ok {
	//			fmt.Println("from comma ok : ", v, ok)
	//			return
	//		} else {
	//			fmt.Println("from comma ok : ", v)
	//			return
	//		}
	//	}
	//}
	go func() {
		for v := range e {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

func send(e, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func SendRecv() {

	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	// send
	go send(eve, odd)
	// recv
	go recv(eve, odd, fanin)
	// fan in values
	for v := range fanin {
		fmt.Println("Value : ", v)
	}

	fmt.Println("about to exit")

}
