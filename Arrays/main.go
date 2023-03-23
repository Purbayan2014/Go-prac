package Arrays

import (
	"fmt"
)

func Arrays() {
	// int array of size 5
	var x [5]int
	x[2] = 1
	fmt.Printf("%d", x[2])

	// grouping data
	y := []int{1, 2, 3, 4, 5}
	fmt.Println(y)

	for indx, val := range y {
		fmt.Printf("\n%d\t%d", indx, val)
	}

	// slicing
	fmt.Println(y[:])   // all elements
	fmt.Println(y[1:3]) // upto 3 from 1 but not including 3
	fmt.Println(y[1:])  // from 1 to the rest of the elements

	// appending
	y = append(y, 6, 7, 8, 9)
	fmt.Println(y)

	z := []int{10, 11, 12, 13}
	y = append(y, z...)

	// deleting using slice
	// lets remove 5
	y = append(y[:4], y[6:]...) // 1,2,3,4 is the first slice and then after appending the rem

	// make
	// has a length of 10 but a capacity of 100 so we can append 100 more
	zy := make([]int, 10, 100) // all will be init 0
	fmt.Println(zy)
	fmt.Println(cap(zy))
	fmt.Println(len(zy))

	// 2d array
	jb := []string{"james", "markus"}
	yb := []string{"jacob", "connie"}
	xp := [][]string{jb, yb}
	fmt.Println(xp)

	// maps or hashmaps in go
	mymap := map[string]int{
		"james":  1,
		"jacob":  2,
		"connie": 3,
	}

	fmt.Println(mymap)
	fmt.Println(mymap["james"])

	delete(mymap, "james") // delete the key in the map

}
