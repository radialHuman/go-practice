package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmplPath string) {

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmplPath)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error while parsing the template %s\n", tmplPath)
		return
	}

}
