package non_ptr_recv_ptr_val 

import (
  "fmt"
  "math"
)

type circle struct {
  radius float64
}

type shape interface {
  area() float64
}

func (c *circle) area() float64 {
  return math.Pi * c.radius * c.radius;
}

func info(s shape) {
  fmt.Println(s.area())
}

func Non_ptr_recv_ptr_val() {

  // now this ptr val takes a ref 
  c := circle{6}
  fmt.Printf("%T\n",&c)
  info(&c)
}
