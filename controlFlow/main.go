package controlFlow

import (
	"fmt"
)

func ControlFlow() {
	// types of for loop
	indx := 1
	for indx <= 3 {
		fmt.Println(indx)
		indx += 1
	}

	for jndx := 7; jndx <= 9; jndx++ {
		fmt.Println(jndx)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// nested loops
	for l1 := 1; l1 <= 3; l1++ {
		for l2 := 1; l2 <= l1; l2++ {
			fmt.Println(l1, l2)
		}
	}

	// printing the ascii
	for indx = 44; indx < 53; indx++ {
		fmt.Printf("\n%v\t%#x\t%#U\t\n", indx, indx, indx)
	}

	// assign conditionally
	if y := 42; y == 42 {
		fmt.Printf("\n%#x", y)
	}

	// else if
	x := 42
	if x == 40 {
		fmt.Println(40)
	} else if x == 41 {
		fmt.Println(41)
	} else {
		fmt.Println(42)
	}

	// switch
	switch {
	case false:
		fmt.Println("false")
	case (2 == 4):
		fmt.Println("false")
	case (3 == 3):
		fmt.Println("true")
	case (4 == 4):
		fmt.Println("true")
	case "s" == "s":
		fmt.Println("s")
	default:
		fmt.Println("boring so default")
	}
}
