package structs

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

// embeddings structs
type agent struct {
	person
	alive bool
}

func Structs() {
	p1 := person{
		fname: "markins",
		lname: "marked",
	}

	fmt.Println(p1)
	fmt.Printf("first name : %v last name : %v", p1.fname, p1.lname)

	a1 := agent{
		person: person{
			fname: "john",
			lname: "rogan",
		},
		alive: false,
	}

	fmt.Println(a1)
	fmt.Printf("\n First name : %v Last Name : %v Alive : %v", a1.fname, a1.lname, a1.alive)

	// Anon structs
	Zoo := struct {
		loc    string
		animal string
	}{
		loc:    "NY",
		animal: "Lion",
	}

	fmt.Println(Zoo)

}
