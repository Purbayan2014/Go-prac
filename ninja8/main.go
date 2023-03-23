package ninja8

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First string
	Age   int
}

type user2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u user) String() {
	fmt.Println(u.First, u.Age)
}

func (u user2) String() {
	fmt.Println(u.First, u.Last, u.Sayings, u.Age)
}

type ByAge []user2

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

type ByLast []user2

func (b ByLast) Len() int {
	//TODO implement me
	return len(b)
}

func (b ByLast) Less(i, j int) bool {
	//TODO implement me
	return b[i].Last < b[j].Last
}

func (b ByLast) Swap(i, j int) {
	//TODO implement me
	b[i], b[j] = b[j], b[i]
}

func Ninja8() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	s := []byte(`[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`)
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}
	users := []user{u1, u2, u3}
	fmt.Printf("%v\n", users)
	ms, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	_, err = os.Stdout.Write(ms)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(s, &users)
	fmt.Println()
	fmt.Println(users)

	fmt.Printf("Before sorting %v\n", xi)
	sort.Ints(xi)
	fmt.Printf("After sorting %v\n", xi)
	fmt.Printf("Before sorting %v\n", xs)
	sort.Strings(xs)
	fmt.Printf("After sorting %v\n", xs)

	u11 := user2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u22 := user2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u33 := user2{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	userss := []user2{u11, u22, u33}
	fmt.Println(userss)
	fmt.Println("By age")
	for _, u := range userss {
		fmt.Println(u.First, u.Last, u.Age, u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
	sort.Sort(ByAge(userss))
	fmt.Println("By sayings")
	for _, u := range userss {
		fmt.Println(u.First, u.Last, u.Sayings)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
	sort.Sort(ByLast(userss))
	fmt.Println("By Last")
	for _, u := range userss {
		fmt.Println(u.First, u.Last, u.Age, u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

}
