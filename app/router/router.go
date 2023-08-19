package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"app/controllers"
)

func StartServer(){
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HomeHandler)
	r.HandleFunc("/about", controllers.AboutHandler)
	r.HandleFunc("/contact", controllers.ContactHandler)

	http.Handle("/", r)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}