// Go-prac package created by markins consists of the all types of
// fundamentals of go in the form of subroutines

package main

import (
	"Go-prac/Applications/applications_brcypt"
	"Go-prac/Applications/applications_json"
	"Go-prac/Applications/applications_sort"
	"Go-prac/Channels"
	"Go-prac/Channels/Context"
	"Go-prac/Channels/FanOut"
	"Go-prac/Channels/send_recv"
	"Go-prac/Concurrency"
	"Go-prac/Concurrency/Atomic"
	"Go-prac/Concurrency/race_condition"
	"Go-prac/ErrorHandling"
	"Go-prac/ErrorHandling/CustomErrorHandling"
	"Go-prac/InputScans"
	"Go-prac/constants"
	"Go-prac/controlFlow"
	fmtpkg "Go-prac/fmtPkg"
	"Go-prac/helloWorld"
	iotacust "Go-prac/iota"
	"Go-prac/methodsets/non_ptr_recv_plus_val"
	"Go-prac/methodsets/non_ptr_recv_ptr_val"
	"Go-prac/ninja1"
	"Go-prac/ninja10"
	"Go-prac/ninja11"
	"Go-prac/ninja2"
	"Go-prac/ninja3"
	"Go-prac/ninja4"
	"Go-prac/ninja6"
	"Go-prac/ninja7"
	"Go-prac/ninja8"
	"Go-prac/ninja9"
	"Go-prac/pointers"
	"Go-prac/shortDeclare"
	stringtype "Go-prac/stringType"
	"Go-prac/structs"
	"Go-prac/varKey"
	"Go-prac/zeroVal"
	"fmt"
)

func main() {
	helloWorld.HelloWorld()
	shortDeclare.ShortDeclare()
	varKey.VarKey()
	zeroVal.ZeroVal()
	fmtpkg.Fmtpkg()
	ninja1.Exer1()
	stringtype.Stringtype()
	constants.Constants()
	iotacust.Iota()
	ninja2.Exer1()
	controlFlow.ControlFlow()
	ninja3.Ninja3()
	ninja4.Exerc4()
	structs.Structs()
	ninja6.Exer6()
	pointers.Pointers()
	non_ptr_recv_plus_val.Non_ptr_recv_plus_val()
	non_ptr_recv_ptr_val.Non_ptr_recv_ptr_val()
	ninja7.Ninja7()
	applications_json.Json()
	applications_sort.Applications_sort()
	applications_brcypt.Bcryptor()
	ninja8.Ninja8()
	Concurrency.Concurrency()
	race_condition.RaceCondition()
	Atomic.RaceCondition()
	ninja9.Ninja9()
	Channels.Channels()
	send_recv.SendRecv()
	FanOut.FanOut()
	Context.Context()
	Context.Context2()
	ninja10.Ninja10()
	InputScans.InputScans()
	// both used panicln so to view them alternately
	// call alternatively by moving these two up and down
	CustomErrorHandling.CustomErrorHandling()
	ErrorHandling.ErrorHandling()
	ninja11, err := Ninja11.Ninja11()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ninja11)

}
