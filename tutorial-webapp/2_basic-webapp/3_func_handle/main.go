package main

import (
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

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Printf("Check out localhost:%s", postNumber)
	_ = http.ListenAndServe(postNumber, nil)

}
