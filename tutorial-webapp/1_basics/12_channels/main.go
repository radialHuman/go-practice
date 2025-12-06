package main

import (
	"fmt"

	"github.com/radiahuman/go-practice/helpers"
)

// #doubt why use channel to pass numbers around? what was wrong with functions params
// #doubt how many values can a chanel handle? if more does it have to be a slice?
const randNumRange = 10

func CalculateValue(ic chan int) { // function that takes in a pram thats channel of int
	// this will do somethign with a random numebr that will be egenretd by a func in helper file
	randomNumber := helpers.RandomNumber(randNumRange) // this will not import untill go mode init is run for this folder
	// now we use the channel to store this random number to be used anywhere else
	ic <- randomNumber
}

func main() {
	// channels are used to send info from one part of the prog to another easily
	// native to go
	intChan := make(chan int) // initalizing a channel which can trasfer only ints to various parts fo the program
	// to make sure that everythign used or opened is closed
	defer close(intChan) // means as soons as this function is closed run things after defer
	// used while reading opening files and db connections

	// now going to call the calculatevalue function using coroutine for simultaneous running of functions
	go CalculateValue(intChan)

	// to get whats coming from channel
	num := <-intChan
	fmt.Println("The numebr is ", num)

}
