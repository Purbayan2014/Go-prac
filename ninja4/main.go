package ninja4

import (
	"fmt"
)

func Exerc4() {
	arr := make([]int, 5, 100)
	indx := 1
	for start := range arr {
		arr[start] = indx
		indx++
		fmt.Println(arr[start])
		fmt.Printf("%T", arr[start])
	}

	slice := arr[:11]
	for indx := 0; indx < len(slice); indx++ {
		slice[indx] = slice[indx] * 2
	}

	for indx := range slice {
		fmt.Printf("\n%d\t%d\t", slice[indx], indx)
	}

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	y := []int{53, 54, 55}
	x = append(x, y...)
	yy := []int{56, 57, 58, 59, 60}
	x = append(x, yy...)

	fmt.Println(x)

	x2x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x2x = append(x2x[:3], x2x[6:]...)
	fmt.Println(x2x)

	c_states := make([]string, 50, 100)
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
                    Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
                    Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
                    Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
                    Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	c_states = append(c_states, states...)
	names := []string{"James", "Bond", "Shaken, not stirred"}
	pronouns := []string{"Miss", "Moneypenny", "Helloooooo", "James."}
	dim_strings := [][]string{names, pronouns}

	for indx, val := range dim_strings {
		fmt.Println("Record : ", indx)
		for jndx, val2 := range val {
			fmt.Printf("\tindex pos : %v \t value : %v", jndx, val2)
		}
		fmt.Println()
	}

	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m {
		fmt.Println("this is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	delete(m, "no_dr")

}
