package main

// html template from server
// created in the templates folder
// tmpl is just the html but in go template

import (
	"fmt"
	"html/template"
	"net/http"
)

// global variable
const postNumber = ":8080" // will notbe changed by any other line of code, its constant else use var

// Home is home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	// need to render go lang template files
	// its static for now
	renderTemplate(w, "home.page.tmpl")
	// any chnages to the tmpl file will work on refreshing of the tab
}

// About is a about page handler
func About(w http.ResponseWriter, r *http.Request) { // #doubt why have this w and r boiler plate , is there anything else these page functions can take in?
	renderTemplate(w, "about.page.tmpl")

}

func renderTemplate(w http.ResponseWriter, tmplPath string) {
	// parse the template
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmplPath)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error while parsing the template %s\n", tmplPath)
		return
	}
	// #doubt even when nothgin i sreturned still the parsed template is shown how?

}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Check out localhost:%s", postNumber)
	_ = http.ListenAndServe(postNumber, nil)

}
