package iota

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

func Iota() {
	fmt.Printf("\n%d\n%d\n%d", a, b, c)
}
