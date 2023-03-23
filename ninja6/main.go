package ninja6

import (
	"fmt"
)

type person struct {
	fname             string
	lname             string
	icecream_flavours []string
}

type collection struct {
	per  person
	per2 person
}

type internal struct {
	doors  int
	colors string
}

type vechile struct {
	inter internal
}

type truck struct {
	v1         vechile
	fourWheeel bool
}

type sedan struct {
	v1     vechile
	luxury bool
}

func Exer6() {
	p1 := person{
		fname:             "john",
		lname:             "wick",
		icecream_flavours: []string{"vanilla", "choco-chip", "caramel"},
	}

	p2 := person{
		fname:             "bella",
		lname:             "swan",
		icecream_flavours: []string{"strawberry", "choco", "choco-fudge"},
	}

	coll := collection{
		per:  p1,
		per2: p2,
	}

	fmt.Printf("first name : %v last name : %v icecream_flavours : %v\n", coll.per.fname, coll.per.lname, coll.per.icecream_flavours)
	fmt.Printf("first name : %v last name : %v icecream_flavours : %v\n", coll.per2.fname, coll.per2.lname, coll.per2.icecream_flavours)

	newmap := map[string][]string{
		p1.lname: p1.icecream_flavours,
		p2.lname: p2.icecream_flavours,
	}

	for indx, val := range newmap {
		fmt.Println("Record : ", indx)
		for jndx, val2 := range val {
			for flav := range val2 {
				fmt.Printf("Key %v flav %v", jndx, val2[flav])
			}
		}
		fmt.Println()
	}

	sed := sedan{
		v1: vechile{
			inter: internal{
				doors:  4,
				colors: "white",
			},
		},
		luxury: true,
	}

	trk := truck{
		v1: vechile{
			inter: internal{
				doors:  4,
				colors: "red",
			},
		},
		fourWheeel: true,
	}

	fmt.Printf("Doors : %v Colors : %v luxury : %T", sed.v1.inter.doors, sed.v1.inter.colors, sed.luxury)
	fmt.Printf("Doors : %v Colors : %v fourWheeel : %T", trk.v1.inter.colors, trk.v1.inter.colors, trk.fourWheeel)

}
