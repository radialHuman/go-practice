package main

import "log"

func main() {
	var aString string // not a global variable still gets affected by the external function, out of scope
	aString = "Something"

	log.Printf("Inital value is %s\n", aString)
	changeStringPointer(&aString) // & is sending the reference to the variable to the function
	log.Printf("new value is %s\n", aString)
}

func changeStringPointer(param *string) { // this is a pointer to a location where a string resides
	log.Printf("Function parameter is %d\n", param) // hexadecimal is also a digit like int #doubt
	newValue := "Nothing"
	*param = newValue
	// no return
}
