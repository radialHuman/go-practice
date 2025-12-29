package main

import (
	"fmt"
	"time"
)

type customer struct {
	fName string
	lName string
	age   int
}

// what if there needs to be a struct inside struct like customer info
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // in built time package, nansec precision
	customer            // need not give it a variable name, this is struct embedding. can be accessed by order.fName
	// no need of order.customer.fName
}

func main() {
	fmt.Println(("======================Struct Embedding============================="))
	orderOne := order{
		id:        "1",
		amount:    15.2,
		status:    "done",
		createdAt: time.Now(),
	} // this will give a struct inside struct with zero default values
	fmt.Println("Struct without struct embedding : ", orderOne)

	orderTwoEmbedding := customer{
		fName: "some",
		lName: "one",
		age:   12,
	}
	orderTwo := order{
		id:        "1",
		amount:    15.2,
		status:    "not done",
		createdAt: time.Now(),
		customer:  orderTwoEmbedding,
	} // this will give a struct inside struct with zero default values
	fmt.Println("Struct with seperate struct embedding : ", orderTwo)

	orderThree := order{
		id:        "1",
		amount:    15.2,
		status:    "not done",
		createdAt: time.Now(),
		customer: customer{
			fName: "some",
			lName: "two",
			age:   10,
		},
	} // this will give a struct inside struct with zero default values
	fmt.Println("Struct with lambda struct embedding : ", orderThree)

	orderThree.customer.age = 30
	fmt.Println("Value changed  : ", orderThree.customer.age)

}
