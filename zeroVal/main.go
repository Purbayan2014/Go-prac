package zeroVal

import (
	"fmt"
)

var ptrvar *string = nil

func ZeroVal() {
	stringvar := "" // zeroval for string
	intvar := 0
	boolvar := false
	floatvar := 0.0

	fmt.Printf("\nNull types")
	fmt.Println(stringvar)
	fmt.Println(intvar)
	fmt.Println(boolvar)
	fmt.Println(floatvar)
	fmt.Println(ptrvar)
}
