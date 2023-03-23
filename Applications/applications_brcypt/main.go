package applications_brcypt

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Bcryptor() {
	s := "pass123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Encrypted with bcrypt")
	fmt.Println(bs)

	fmt.Println("Compare the hashed password with the original password")
	err = bcrypt.CompareHashAndPassword(bs, []byte(s))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hashed password is right")
}
