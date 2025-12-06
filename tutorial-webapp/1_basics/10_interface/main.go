package main

import "fmt"

// doubt how to import functions from other fodlers

// defining an interface
type NameOfInterface interface {
	// signature : list of fucntion that an interface should have
	Function1() string  // this inetrface must have this function that takes no param and returns string
	Function2(int) bool // same but take int and gives bool
}

type InterfaceType1 struct {
	Name string
	Age  int
}

type InterfaceType2 struct {
	Name      string
	Available bool
}

func main() {
	/*
	   Interface is a contract
	   is a bunch of rules for types that follow the interface
	*/

	person := InterfaceType1{Name: "Name1", Age: 20}
	tool := InterfaceType2{Name: "Hammer", Available: true}
	Printer(person) // thoguh perinter accepst param of type nameofinstance but it was okay with interfacetype 1 because it had the essential functions associated with it
	// doubt did nto understand the need for interface
	Printer(tool) // this will tn work if fucntions are not associated with it
}

// creating functions for the structs (receivers) which will satisfy interface's requriement
func (p InterfaceType1) Function1() string {
	return p.Name
}

func (p InterfaceType1) Function2(int) bool {
	return p.Age > 30
}

func (p InterfaceType2) Function1() string {
	return p.Name
}

func (p InterfaceType2) Function2(int) bool {
	return p.Available
}

// now using the intefrcae
func Printer(i NameOfInterface) {
	fmt.Println(i.Function1(), i.Function2(40))
}
