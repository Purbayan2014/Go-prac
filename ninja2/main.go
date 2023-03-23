package ninja2

import (
	"fmt"
)

const (
	fyrs   = iota * 10
	syrs   = iota * 20
	thyrs  = iota * 30
	fouyrs = iota * 40
)

func Exer1() {
	y := 45
	fmt.Printf("\n%d\t%b\t%#X\n", y, y, y)
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println(a, b, c, d, e, f)
	newcar := 44
	fmt.Printf("\n%d\t%#X\t%b", newcar, newcar, newcar)
	newvar := newcar << 1
	fmt.Printf("\n%d\t%#X\t%b", newvar, newvar, newvar)
	fmt.Printf("\n\t%d\t%d\t%d\t%d\n", fyrs, syrs, thyrs, fouyrs)
}
