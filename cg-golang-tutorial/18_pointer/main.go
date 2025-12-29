package main

import "fmt"

/*
POINTERS
its the memory locaiton address of any variable

*/

// by value
func changeInt(a int) { // in a function when a variable is passed, a copy of that number is passed not the actual location
	a = 5
	fmt.Println("In chnageInt : ", a)
}

// by reference
func changeIntProper(a *int) { // pointer
	*a = 5 // dereferencing, else it will not take the value in
	fmt.Println("Address of num in proper function : ", &a)
	fmt.Println("In chnageIntProperly : ", *a)

}

func main() {
	fmt.Println(("======================Pointers============================="))
	num := 10
	changeInt(num) // though we have asked to change the value of this it changed in the function but not in main
	// its apssing the  variable to the function by value not reference
	fmt.Println("In Main : ", num)
	fmt.Println("Address of num in Main : ", &num)
	fmt.Println(("======================Proper way============================="))
	changeIntProper(&num) // pass with reference, get as pointer, dereference it
	fmt.Println("In Main after proper reference: ", num)

}
