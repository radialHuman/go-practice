package main // need not be main, it can be anything but its a common thing so keep it main

import (
	"fmt"
	"log"
)

func main() { // this gets executed first
	var functionsOutput string               // no space, no starting with numbers. usually use camel case, strongly typed
	functionsOutput = function1("something") // unsuaed variables are error and vscode highlights if not sued
	// at this point though the variable is assigned but not really helpful it will still be an error
	fmt.Println((functionsOutput)) // doesn seem to affect the output even with double brackets #doubt
	// another way of printing, with dates and all like logging
	log.Println(functionsOutput)

	// if storing is not required then function can directly be passed as in other languages
	log.Println(function1("else"))

	// ****************************************************************************************
	var number int  // another type
	println(number) // if this works why use fmt? #doubt
	// numbers if uninitialized has the value 0 as default

	number = 7
	number = 8
	fmt.Printf("%d\n", number) // like using f string or format

	// ****************************************************************************************
	var secondFunctionOutput string
	functionsOutput, secondFunctionOutput = function2("double variable storing") // output of funciton with n variables must be stored in n variables
	fmt.Printf("First word is %s\nSecond word is %s\n", functionsOutput, secondFunctionOutput)

	functionsOutput, _ = function2("double, but ignoring the other\n") // use _ to ignore the in coming output from function
	fmt.Printf("First word is %s", functionsOutput)

}

// apart from main function there can be other functions, but main is default

func function1(param1 string) string { // string input as param1 and string output, no : no ->
	return param1
}

func function2(param1 string) (string, string) { // multiple return
	return param1, " ****"
}
