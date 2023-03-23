package non_ptr_recv_plus_val

import (
	"fmt"
	"math"
)

// circle type
type circle struct {
	radius float64
}

// interface with prop area
type shape interface {
	area() float64
}

// (c circle) is the method set which is used to implement the area functions
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// info function which uses shape to print the area
func info(s shape) {
	fmt.Println("area", s.area())
}

func Non_ptr_recv_plus_val() {
	// create a circle
	c := circle{
		radius: 22,
	}
	info(c)
}
