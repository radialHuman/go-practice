package main

import "fmt"

func main() {
	age := 17.9
	if age >= 18 { //if
		fmt.Println("Adult")
	} else if age <= 5 { // else if
		fmt.Println("Kid")
	} else { // else
		fmt.Println("Minor")
	}

	var user string = "admin"
	var hasPermission = true
	if user == "admin" || hasPermission { // or
		fmt.Println("Allowed")
	}
	user = "user"
	hasPermission = false
	if user != "admin" && !hasPermission { // and
		fmt.Println("Not Allowed")
	}

	// variable declaration inside if
	if age := 1; age >= 18 {
		fmt.Println("Adult : ", age)
	} else if age < 5 {
		fmt.Println("Kid : ", age)

	}

	// no ternary operator  ? : in go, use if else
	// prioritize readability and explicitness over compact syntax.
}
