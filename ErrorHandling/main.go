// Package ErrorHandling Package for error handling
package ErrorHandling

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered r in f ", r)
		}
	}()
	fmt.Println("calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
}

func ErrorHandling() {

	f()
	fmt.Println("Returned normally from f")
	file, err := os.Create("Names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fs, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(fs))

	logfile, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	file2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)

	}
	defer file2.Close()

	_, err = os.Open("no-file1.txt")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = os.Open("no-file2.txt")
	if err != nil {
		log.Panicln(err)

	}

	fmt.Println("check the logs file in the directory")

	//os.Exit(1) // exists the code

}
