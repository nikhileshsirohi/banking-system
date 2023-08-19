package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikhileshsirohi/banking-system/app/controllers"
	"github.com/nikhileshsirohi/banking-system/app/repo"
	"github.com/nikhileshsirohi/banking-system/app/router/middlewares"
)

func StartServer() {
	repo.StartDB()

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HomeHandler)
	r.HandleFunc("/about", controllers.AboutHandler).Methods("POST")
	r.HandleFunc("/contact", controllers.ContactHandler)

	registerUserRoute(r)

	http.Handle("/", r)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func registerUserRoute(r *mux.Router) {
	routeGroup := r.PathPrefix("/v1/user").Subrouter()
	routeGroup.Use(middlewares.BasicAuthMiddleware)

	routeGroup.HandleFunc("/signup", controllers.Signup).Methods("POST")
}
