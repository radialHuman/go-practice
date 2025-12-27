package main

import (
	"fmt"
	"time"
)

func main() {
	// too many if else can be avoided
	var age int = 7

	// simple switch
	println("=================Simple switch=====================")
	switch age {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	case 6:
		fmt.Println(6)
	default:
		fmt.Println(-1)
	}
	println("============Conditional switch==============")
	switch time.Now().Weekday() { // check todays date
	case time.Saturday, time.Sunday: // or
		fmt.Println("Weekend")
	default:
		fmt.Println("Go work")
	}

	println("============Type switch==============")
	// lambda function, functiosn are first class, so they can be passed around
	// lambdaFunction := func(i interface{}) { // inetrface is anytime accepted, old way of doing it
	lambdaFunction := func(i any) { // inetrface is anytime accepted #new go 1.18
		switch inputParameter := i.(type) { // #new
		case int:
			fmt.Println("Integer")
		case float64:
			fmt.Println("Decimal")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Something else", inputParameter)
		}
	}
	lambdaFunction("here")
	lambdaFunction(1)
	lambdaFunction(2.0) // taken as 64 by default, i guess
	lambdaFunction(true)
}
