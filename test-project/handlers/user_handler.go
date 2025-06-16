package handlers

import (
	"indal/services"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	newUserService := services.NewUserService()
	err := newUserService.CreateUser(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error creating user"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created successfully"))
}
