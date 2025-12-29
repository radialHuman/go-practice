package main

import (
	"fmt"
	"time"
)

/*
STRUCTS
GO doesn have classes so use structs
custom data structure, mix
*/

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // in built time package, nansec precision
} // struct structure/blueprint declared
// #doubt can structs have functions like classes? LOOK DOWN

// receiver type #new
// to connect the fucntion to the above struct
// add a variable before the function name with first letter as that of the struct name and its type
func (o order) changeStatus(a string) {
	o.status = a // all the elements of the struct will be avilable to this fucntion
	fmt.Println("Struct inside function : ", o.status)
} /// this will not work, as its not using pointer and a copy fo the struct is passed

func (o *order) changeStatusProperly(a string) { // now that the pointer to the struct is passed it will work
	// #doubt no need to deferencing or sending refernce while function call?
	o.status = a // all the elements of the struct will be avilable to this fucntion
	fmt.Println("Struct inside proper function : ", o.status)
}

func (o order) getAmount() float32 { // since nthing is being changed, poitner is not required
	return o.amount
}

// CONSTRUCTORS #new
func newOrder(id string, amount float32, status string) *order { // the convention is to use new followed by the struct name
	// go doesn allow  countructro by defualt
	// NewOrder will amke ti accessible to other files, newOrder will keep things bound to this file, applicable to the avriables also
	/* USES of constructors
	- to initialize structs in a known, valid state, enhancing code maintainability and readability
	- go encourages the use of factory functions to achieve similar functionality
	- ensuring that all necessary initialization steps are performed consistently.
	- ex : a constructor can validate input parameters, set sensible default values, or initialize complex internal state, preventing the creation of invalid or inconsistent objects
	- handle optional parameters through the functional options pattern, manage complex construction via builder patterns, or return interfaces to hide implementation details while exposing only necessary methods.
	- promotes loose coupling and better abstraction, especially when working with dependencies modeled as interfaces.
	- perform setup tasks such as establishing connections, initializing caches, or validating dependencies, which may result in errors that should be propagated to the caller
	- considered a best practice for creating clean, reliable, and maintainable code, especially when the zero value of a struct is insufficient or when initialization logic is non-trivial

	*/
	orders := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &orders // since the values are changing, the pointer and reference needs to be sent
	// needs to be deferenced when used
}

func main() {
	fmt.Println(("======================Struct============================="))
	// creatng instance of struct just like classes
	orderOne := order{
		id:     "a",
		amount: 50.0,
		status: "done",
		// createdAt:, // even if not given its fine, it gets the default zero value
	}
	fmt.Println("Struct : ", orderOne)
	fmt.Println("Struct element : ", orderOne.status)

	// adding an element later
	orderOne.createdAt = time.Now()
	fmt.Println("Struct after adding : ", orderOne)

	fmt.Println(("======================Second order============================="))

	orderTwo := order{
		id:        "b",
		amount:    10.0,
		status:    "not done",
		createdAt: time.Now(),
	}

	fmt.Println("Another Struct : ", orderTwo)

	fmt.Println(("======================Change status using attached function============================="))
	fmt.Println("Struct before change in status : ", orderOne.status)
	orderOne.changeStatus("Not done")
	fmt.Println("Struct after changing status using attacked function: ", orderOne.status)

	fmt.Println(("======================Change status using attached proper function============================="))
	fmt.Println("Struct before change in status : ", orderOne.status)
	orderOne.changeStatusProperly("Not done")
	fmt.Println("Struct after changing status using attacked function: ", orderOne.status)

	fmt.Println(("====================== function wihtout pointer============================="))
	fmt.Println("Amount is : ", orderOne.getAmount())
	fmt.Println("Another amount is : ", orderTwo.getAmount())

	fmt.Println(("====================== strcut creation wiht constructor============================="))
	orderThree := *newOrder("3", 0.0, "unknown") // since the retuned value is a reference, it has to be dereferenced here
	// but it will work even if not dereferenced, as struct takes care of that
	fmt.Println("Consructor Struct : ", orderThree)

	fmt.Println(("====================== one time struct, lambda type ============================="))
	orderFour := struct {
		id     string
		amount float32
		status bool
	}{
		"4", 12.0, true,
	}
	fmt.Println("Lambda Struct : ", orderFour)

}
