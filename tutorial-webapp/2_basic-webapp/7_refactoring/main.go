package main

// hacing all the funcitons in one file makes it unreadable and difficult to maintain
// so splitting it up into myutiple ones is a good idea
// having logically related functions must be in the same file
// creating a folder handlers
// if just main.go is now run it will not work as it will not knwo about the handler.go file
// so we need to run them all and since home and baout are public functions,
// it will find them in all the other go files and use that to connect
// go run *.go

// for html page formatting using bootstrap css via link in head in tml, to make it look good,
// #doubt, its cdn delivered, no static css file created, cant work wihtout net?
// using div and intellisens for html to make text in breowser to not be left aligned

import (
	"fmt"
	"net/http"
)

const postNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Check out localhost:%s", postNumber)
	_ = http.ListenAndServe(postNumber, nil)

}
