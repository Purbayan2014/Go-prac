package ninja9

import (
	"fmt"
	"runtime"
	"sync"
)

type person struct {
	voice string
}

type human interface {
	speak() string
}

func (p *person) speak() string {
	return fmt.Sprintf("%s", p.voice)
}

func saySomething(hum human) {
	fmt.Println(hum.speak())
}

var wait sync.WaitGroup

func test1() {
	fmt.Println("test1 subroutine")
	wait.Done()

}

func test2() {
	fmt.Println("test2 subroutine")
	wait.Done()
}

func Ninja9() {

	wait.Add(1)
	go test1()
	wait.Add(1)
	go test2()

	// create a person
	p1 := person{"hello from markins"}
	saySomething(&p1)

	// implementing the race condition
	fmt.Println("Ninja 9 race condition")
	fmt.Println("CPU : ", runtime.NumCPU())
	counter := 0
	const gs = 100
	var mutex sync.Mutex
	wait.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mutex.Unlock()
			wait.Done()
		}()
		fmt.Println("Goroutines : ", runtime.NumGoroutine())
	}
	wait.Wait()
	fmt.Println("Ninja 9 race condition ends")
	fmt.Println("OS : ", runtime.GOOS)
	fmt.Println("arch : ", runtime.GOARCH)

}
