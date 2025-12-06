package main

import (
	"fmt"
	"time"
)

// creating a bunch of variables to store records cna be done in the old fashioned way
/*
var fName string
var lName string
var dob string
*/
// passing them in a function will make it unreadable
// what if there are other functions hacing the same variabl names, it will be confusing

// but a better way is using struct
// things thar are in capital camel case, can be called from outside the packge or scope #doubt
type User struct { // its like a kv pair, and its vairbales are not in camel case
	FirstName string
	LastName  string
	Dob       time.Time // a in built way of date type, using string is risky
}

func main() {
	// initializing the struct
	user := User{
		FirstName: "John",
		LastName:  "Doe",
		// Dob: ,
	}

	fmt.Printf("The first name is %s\n", user.FirstName)
	fmt.Printf("The dob is %s\n", user.Dob) // though unintialted, the value is a default time

}
