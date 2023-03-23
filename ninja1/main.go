package ninja1

import (
	"fmt"
)

// Exer2
var x1 int
var y1 string
var z bool

// Exer3
type hotdog int

var newx hotdog
var newy int

func Exer1() {
	x := 42
	y := "James Bond"
	st := true

	fmt.Printf("\n%d\t%s\t%t", x, y, st)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(st)

	x1 = 22
	y1 = "qweqwe"
	z = true

	fmt.Printf("\n%d\t%s\t%t", x1, y1, z)

	newx = 22
	newy = int(newx)
	fmt.Println(newx)
	fmt.Println(newy)
}
