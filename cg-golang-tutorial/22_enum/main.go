package main

import "fmt"

/*
Enum
In case a thing can have multiple value which are at different places, and are passed as string
this can lead to typos and errors.
Instead like db normalization use enum which will give ids and use them instead

Enums are not in built in go
made using type and const


Need to make custom datatype first
*/

type OrderType int // from now on ordertype is alias for int

const (
	Received OrderType = iota // iota starts with 0 and keeps increasing by 1 for further consts
	Confirmed
	Pending
	Cancelled
)

// another enum but for string type
type OrderName string

const (
	received  OrderName = "r" // for string it has to be declared for each case
	confirmed OrderName = "co"
	pending   OrderName = "p"
	cancelled OrderName = "ca"
)

func ChangeStatus(status OrderType) {
	fmt.Printf("The status numbers : %d\n", status)
}

func ChangeName(status OrderName) {
	fmt.Printf("The status name : %s\n", status)
}

func main() {
	fmt.Println(("======================Int Enum============================="))
	ChangeStatus(Received)  // not a string , is a const so less chance of error
	ChangeStatus(Confirmed) // more readable, easy to type
	ChangeStatus(Pending)
	ChangeStatus(Cancelled)

	fmt.Println(("======================String Enum============================="))
	ChangeName(received)  // not a string , is a const so less chance of error
	ChangeName(confirmed) // more readable, easy to type
	ChangeName(pending)
	ChangeName(cancelled)

}
