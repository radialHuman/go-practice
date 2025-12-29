package main

import "fmt"

/*
CLOUSER #new
a function object that retains access to variables from its enclosing (lexical) scope, even after the outer function has finished executing.
This mechanism enables the creation of functions with persistent state, useful for tasks like creating function factories, implementing callbacks, or maintaining counters without relying on global variables.
*/

// if a variable is used from the outer scope and in an inner scope, even after calling the function, the values in the function
// will not get deleted
// #doubt not sure what the use is fo this weird behaviours

func counter() func() int { // this functionretruns a function that returns  an integerer
	var count int //default vakye is 0

	return func() int { // lambda function
		count += 1
		return count
	}
}

func main() {
	fmt.Println(("======================Clousers============================="))
	increment := counter()
	fmt.Println(increment()) // should give 1
	fmt.Println(increment()) // should also give 1, but gives 2 due to clouser
	fmt.Println(increment()) // gives 3

}
