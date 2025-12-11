package main

import (
	"errors"
	"fmt"
	"net/http"
)

// global variable
const postNumber = ":8080" // will notbe changed by any other line of code, its constant else use var

// writting it in a better way from pervious version
// having two pages in web  app

// comments for fucntions should start with function name : go style

// Home is home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Landing page")
}

// About is a about page handler
func About(w http.ResponseWriter, r *http.Request) { // #doubt why have this w and r boiler plate , is there anything else these page functions can take in?
	fmt.Fprintf(w, "About page")
}

// Divide is to divide numbers and show
func Divide(w http.ResponseWriter, r *http.Request) {
	output, err := divideFunction(1, 0)
	if err != nil {
		fmt.Fprintf(w, "Cant divide by zero")
		return // this means it will stop else it will go ahead witht he rest and show 0 which cant be right division
	}
	fmt.Fprintf(w, "%f", output)

}

func divideFunction(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("Denominator is 0")
		return 0, err
	} else {
		result := x / y
		return result, nil
	}

}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Check out localhost:%s", postNumber)
	_ = http.ListenAndServe(postNumber, nil)

}
