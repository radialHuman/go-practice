package handlers

import (
	"net/http"

	"github.com/radialhuman/8_go_modules/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.html")

}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")

}
