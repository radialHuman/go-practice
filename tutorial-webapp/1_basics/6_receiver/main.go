package main

import "log"

// struct can have functions associated with it
type aStruct struct {
	Param string
}

func main() {
	var variable aStruct
	variable.Param = "Something"

	// can also be written as
	anotherVariable := aStruct{
		Param: "anything",
	}

	log.Printf("FIrst one is %s\n", variable) // this works for now because the struct has ntohign but string, as sson as there are mixed types in struct this will not work
	log.Printf("Second one is %s\n", anotherVariable)

	variable.printParams() // receiver can be accessed from here but not as a normal function call
}

// function thats typed to a struct is a receiver
func (pr *aStruct) printParams() { // adding the pointer to the struct that ti can access
	log.Printf("Receiver says %s\n", pr.Param)
}
