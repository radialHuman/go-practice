package main

import "fmt"

/*
INTERFACE
is a contract
Convention : it should end with -er

Its an abstraction where anythign can be connected to the app shares common properies defiend in interface
ex : app can connect to any db, an  interface allows to not swap the db when it changes

Using this explanation here : https://youtu.be/SX1gT5A9H-U

// all the methods in the interface must be made avialabe using struct method. even if one is missing, it will nto work
*/

// lets say we have shapes and we need to find some of its proerties like area and circumfrence
type shape interface { // must have methods in it, now replaced with any #doubt, still not sure what ti does
	// juts declare the functions this inerface cna use
	area() float32
	circumfrence() float32
}

// now the shpes can be defined as structs
type rectangle struct {
	length, height float32
}

type circle struct {
	radius float32
}

// so far the structs are not aware of the connection b/w the shape and interface funcitons
// for making the connection make a functionw ith the same name as in inetrface function for each struct

func (r rectangle) area() float32 {
	return r.height * r.length
}

func (r rectangle) circumfrence() float32 {
	return 2 * (r.height + r.length)
}

func (c circle) circumfrence() float32 {
	return 2 * 3.14 * c.radius
}

func (c circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

func findArea(s shape) float32 {
	return s.area()
}

func findCircumfrence(s shape) float32 {
	return s.circumfrence()
}

// function that cantake in any type
func processAnything(t any) any {
	fmt.Printf("Type : %T, value : %v\n", t, t)
	return t
}

func main() {
	fmt.Println(("======================Without interface============================="))
	rect := rectangle{length: 10, height: 12}
	fmt.Println("Area of rectangle : ", rect.area())
	fmt.Println("Circumfrence of rectangle : ", rect.circumfrence())

	cir := circle{radius: 10}
	fmt.Println("Area of circle : ", cir.area())
	fmt.Println("Circumfrence of circle : ", cir.circumfrence())

	fmt.Println(("======================With interface============================="))
	// instea dof calling area and circumfrence for each type we can have a function that will take in the geenric type and
	// irrespective of rectangle or cirlce gives area
	fmt.Println("Area of rectangle : ", findArea(rect))
	fmt.Println("Circumfrence of rectangle : ", findCircumfrence(rect))
	fmt.Println("Area of circle : ", findArea(cir))
	fmt.Println("Circumfrence of circle : ", findCircumfrence(cir))

	fmt.Println(("======================interface type, or any type============================="))
	anyValueOne := 10
	anyValueTwo := true
	processAnything(anyValueOne)
	processAnything(anyValueTwo)

}
