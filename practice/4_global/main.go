package main

import "log"

// global variable : will be used unless a fucntion has param of the same name
// var globeVar string
var globeVar = "String" // automatic type understanding

func main() {
	globeVar := 10 // though earlier string, not an int for this scope. := worked only because it was declared earlier
	// #doubt := cna be used before initializing?
	log.Printf("GLobal variable : %d", globeVar) // log.println is a variadic function, i.e. can take multiple params 0,1,2...
	checkGlobalVar()
}

func checkGlobalVar() {
	log.Printf("Here the global vairbale is %s", globeVar)
}
