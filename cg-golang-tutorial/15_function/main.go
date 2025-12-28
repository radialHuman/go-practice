package main

import "fmt"

// func add(a int, b int) int {
func add(a, b int) int { // both and b are of type int, shorthand
	return a + b
}

func getmultipleReturn() (string, int, string) { // multiple return with each its own type
	return "1", 2, "3"
}

// generally in golang fucntions return the actual value to be returnd and an error, later

func funcInception(func1 func(a int) int) int {
	return func1(1)
}

func main() {
	fmt.Printf("The addtion of 1 and 2 is : %d\n", add(1, 2))
	one, two, three := getmultipleReturn()
	// one, _, three := getmultipleReturn() // to suppress a value while getting
	fmt.Println("Values from getmultipleReturn : ", one, two, three)

	// lambda funciton
	returnNumber := func(a int) int {
		return 2
	}
	fmt.Println(funcInception(returnNumber)) // #doubt did not get the process of calling properly
}
