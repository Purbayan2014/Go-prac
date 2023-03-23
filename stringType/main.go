package stringtype

import (
	"fmt"
)

func Stringtype() {
	s := "My life is not negotiable"
	fmt.Println(s)
	fmt.Printf("\n%T", s)

	// []byte(s) ==> gives the byte slice of the string
	byteSlice := []byte(s)
	fmt.Printf("\n%T", byteSlice)
	fmt.Println(byteSlice)

	// check for utf-8
	for indx := 0; indx < len(s); indx++ {
		fmt.Printf("%#U ", s[indx])
	}
	fmt.Println()
	// index value
	for index, value := range s {
		fmt.Printf("at index position %d we have hex val %#x\n", index, value)
	}
}
