package Ninja11

import (
	"encoding/json"
	"errors"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
	err     error
}

type customError struct {
	err error
}

func (ce customError) Foo() string {
	return fmt.Sprintf("Error : %v", ce.err)
}

func (p person) Error() string {
	return fmt.Sprintf("Error : %v", p.err)
}

func Ninja11() (string, error) {
	p1 := person{
		First:   "james",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
		err:     errors.New("some error"),
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		nme := fmt.Errorf("error :: %v", err)
		return "", person{
			First:   p1.First,
			Last:    p1.Last,
			Sayings: p1.Sayings,
			err:     nme,
		}
	}

	return string(bs), nil
}
