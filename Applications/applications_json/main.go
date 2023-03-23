package applications_json

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

type Animal struct {
	Name  string
	Order string
}

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func Marshalling(gr ColorGroup) {
	fmt.Println("Marshall")
	b, err := json.Marshal(gr)
	if err != nil {
		fmt.Println("error : ", err)
	}
	_, err = os.Stdout.Write(b)
	if err != nil {
		fmt.Println(err)
	}
}

func Json() {
	// creating the group
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"crimson", "orange", "red", "blue"},
	}

	Marshalling(group)
	fmt.Println("Unmarshalling")
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nAll of the data", people)
	for index, val := range people {
		fmt.Println("\nPerson Number : ", index)
		fmt.Println(val.First, val.Last, val.Age)
	}
}
