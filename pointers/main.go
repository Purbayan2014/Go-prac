package pointers

import (
	"fmt"
)


func foo(y *int) {
	fmt.Println("Before deref");
	fmt.Println(y);
	fmt.Println(*y);
	// changing the value 
	*y = 323;
	fmt.Println("After refererencing");
	fmt.Println(y)
	fmt.Println(*y);
}

func Pointers() {
	// storing the address using the &
	a := 22
	b := &a

	// types
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	fmt.Printf("%T\n", b)

	// to get the value back from b
	fmt.Println(*b)
	fmt.Println(*&a) // same way

	// changing the value of b
	*b = 23
	fmt.Println(*b)
	fmt.Println(b);

	// passing a pointer to the function
	c := 223;
	foo(&c);
}
