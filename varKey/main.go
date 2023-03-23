package varKey

import (
	"fmt"
)

// declare a variable outside the function only using var
var newvar string = "markins"

func VarKey() {
	// := only used in the function
	fmt.Printf("\nMy name is %s", newvar)
}
