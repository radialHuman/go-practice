package main

import (
	"log"
	"time"
)

var global = 1 // global variable

// grouping of variables so that its easy to play around and write is a structure

type User struct { // capital means its callable outside this package
	firstName string
	lastName  string
	age       int
	dob       time.Time
}

func main() {
	log.Println("Before local global variable declaration", global) // since no local till now, it will go for global
	global := 2
	log.Println("After local global variable declaration", global) // looks at local first, and use that

	var person1 User
	person1.age = 10 // this is one way
	person1.firstName = "first"
	person1.lastName = "last"
	// the other way is
	log.Println(person1)
	log.Println(person1.firstName)

	person2 := User{
		firstName: "some",
		lastName:  "one",
	}
	log.Println(person2)
	log.Println(person2.age) // not there but gets default value

}

/*OUTPUT
2026/04/04 18:49:11 Before local global variable declaration 1
2026/04/04 18:49:11 After local global variable declaration 2
2026/04/04 18:49:11 {first last 10 {0 0 <nil>}}
2026/04/04 18:49:11 {some one 0 {0 0 <nil>}}
*/
//Building Modern Web Applications with Go (Golang)/2. Overview of the Go Language
