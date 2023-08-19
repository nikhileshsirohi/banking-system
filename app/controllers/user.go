package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nikhileshsirohi/banking-system/app/models"
	"github.com/nikhileshsirohi/banking-system/app/services"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	res, err := services.Signup(user)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Nikhil", "Sirohi")
	w.WriteHeader(GetStatusCode(err))
	json.NewEncoder(w).Encode(res)
	json.NewEncoder(w).Encode(err)

}

func GetStatusCode(err error) int {
	if err == nil {
		return 200
	}
	return 400
}
