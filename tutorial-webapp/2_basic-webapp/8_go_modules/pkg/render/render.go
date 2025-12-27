package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate
func RenderTemplate(w http.ResponseWriter, tmplPath string) {

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmplPath)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error while parsing the template %s\n", tmplPath)
		return
	}

}
