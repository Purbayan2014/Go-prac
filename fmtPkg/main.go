package fmtpkg

import (
	"fmt"
)

func Fmtpkg() {
	y := 42
	fmt.Printf("\n%#x \t %b \t %x", y, y, y)
	// Sprintf assigns value to string
	// Fprintf assigns value to file
}
