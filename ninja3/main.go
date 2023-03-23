package ninja3

import (
	"fmt"
)

var start int32

type favsport string

var sport favsport

func Ninja3() {

	start = 65
	for start <= 90 {
		fmt.Println(start)
		for j := 0; j <= 3; j++ {
			fmt.Printf("\t%#U", start)
		}
		println()
		start++
	}

	sport = "none"
	switch sport {
	case "ice-hockey":
		println("ice-hockey")
	case "soccer":
		println("soccer")
	default:
		println("lazy ass !!")
	}

}
