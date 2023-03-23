package ninja7 

import (
  "fmt"

)


type person struct {
  name string
}

func changeme(p *person) {
  (*p).name = "good name"; 
}

func Ninja7() {
  a := 333;
  b := &a;
  fmt.Println(b);
  p1 := person {
    name : "james",
  }
  fmt.Println(p1)
  changeme(&p1);
  fmt.Println(p1)

}
