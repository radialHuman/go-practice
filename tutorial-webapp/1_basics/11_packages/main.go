package main

import (
	"fmt"

	"github.com/radialhuman/go-practice/helpers"
)

func main() {
	// to create a package : go mod init github.com/radialhuman/go-practice
	// create helpers folder in the same fodler and have a helper.go in it
	var myVar helpers.AStruct // wrirting this will type the import automatically
	fmt.Println(myVar)
	myInt := helpers.AddOne(myVar.AnInt)
	fmt.Println(myInt)

}
