package applications_sort

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// this function tells how to print the struct
func (p Person) String() string {
	return fmt.Sprintf("%s : %d", p.Name, p.Age)
}

// ByAge implements the sort.Interface for the []PERSON based on the age field
type ByAge []Person

func (b ByAge) Len() int {
	//TODO implement me
	return len(b)
}

func (b ByAge) Less(i, j int) bool {
	//TODO implement me
	return b[i].Age < b[j].Age
}

func (b ByAge) Swap(i, j int) {
	//TODO implement me
	b[i], b[j] = b[j], b[i]

}

func Applications_sort() {
	arr := []int{3, 5, 7, 1, 2}
	fmt.Printf("Unsorted Int array : %v\n", arr)
	sort.Ints(arr)
	fmt.Printf("Int Array sorted : %v\n", arr)

	strarr := []string{"james", "jacob", "aristotle"}
	fmt.Printf("Unsorted String array : %v\n", strarr)
	sort.Strings(strarr)
	fmt.Printf("String Array sorted : %v\n", strarr)

	fmt.Println("custom sort based on the age")
	people := []Person{
		{"Bob", 31},
		{"john", 22},
		{"markins", 21},
		{"jenny", 2},
	}

	fmt.Println("Unsorted people")
	fmt.Println(people)
	fmt.Println("sorted based on their ages")
	sort.Sort(ByAge(people))
	fmt.Println(people)

}
