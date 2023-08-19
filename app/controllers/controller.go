package controllers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the About Page.")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("contact_handler_info_receiver")
	fmt.Fprintln(w, "You can contact us at contact@example.com")
}