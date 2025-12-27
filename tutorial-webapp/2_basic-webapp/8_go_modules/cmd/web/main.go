package main

// rearrign things for go modules
// go modules are package mgmt for go
// also passing data from fe to be

// go mod init github.com/radialhuman/<pwd> : makes your folder ready for publishing, makes a go.mod
// go.mod will rarley be edited by hand, commands will populate it

// other files are put in thier respective folders and main is in cmd/web/
// other files are more calling package mian, instead package <their pwd>

// now to run this file it cant be go run *.go anymore
// go run cmd/web/*.go

import (
	"fmt"
	"net/http"

	"github.com/radialhuman/8_go_modules/pkg/handlers"
)

const postNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Check out localhost:%s", postNumber)
	_ = http.ListenAndServe(postNumber, nil)

}
