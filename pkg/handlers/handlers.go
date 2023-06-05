package handlers

import (
	"net/http"

	"github.com/cloudy4next/hotel_reservation_go/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.Rendetemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.Rendetemplate(w, "about.page.html")
}
