package main

import "fmt"

func main() {
	var isTrue bool // boolean

	isTrue = !true

	// if isTrue {// also can be written as
	if isTrue == true {

		fmt.Println("The value of isTreu is", isTrue)
	} else {
		fmt.Println("The value of isTrue is", isTrue)
	}

	//**************************************************
	num := 10
	if isTrue && num == 10 { // || is or
		fmt.Printf("%t : %d\n", isTrue, num) // f string for boolean is %t for true
	} else { // else if ...
		fmt.Printf("%t\n", isTrue)
	}
	//**************************************************
	// myVar := 10
	switch myVar := 10; { // first match will last match
	case myVar < 5:
		fmt.Printf("Less than 20\n")
	case myVar < 10:
		fmt.Printf("Less than 10\n")
	default:
		fmt.Printf("%d\n", myVar)
	}

}
