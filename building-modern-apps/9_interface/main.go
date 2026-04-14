package main

import "log"

type Animal interface {
	// functions that Animal type must have
	// its a contract that must be satisfied to allow structs with these functions to be passed
	// where functions are expecting an interface
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name  string
	Color string
	Age   int
}

func main() {
	dog := Dog{
		Name:  "dog1",
		Breed: "breed1",
	}

	// gorr := Gorilla{
	// 	Name:  "gorr1",
	// 	Color: "color1",
	// 	Age:   10,
	// }

	PrintAnimalDetails(dog) // the function needs a interface but got a struct
	// it accepts because the strcut has 2 functions it required associated with it
	// if it was not satisfying then it would say
	// cannot use dog (variable of struct type Dog) as Animal value in argument to PrintAnimalDetails:
	// Dog does not implement Animal (missing method NumberOfLegs)

}

// adding funtions for structr variables
func (d Dog) Says() string {
	return "hello"
}

func (d Dog) NumberOfLegs() int {
	return 5
}

func PrintAnimalDetails(a Animal) {
	log.Println(a.Says(), a.NumberOfLegs())
}

/*INFO
https://stackoverflow.com/questions/39092925/why-are-interfaces-needed-in-golang

Interfaces are a tool. Whether you use them or not is up to you, but they can make code clearer, shorter, more readable, and they can provide a nice API between packages, or clients (users) and servers (providers).

Not all packages will benefit from them, but for certain programming tasks interfaces can be an extremely useful for abstraction and creating package APIs,
*/
