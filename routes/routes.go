package routes

import (
	"net/http"

	"github.com/D-Toshchakov/pet/social-media/handlers"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {
	r.HandleFunc("/health", handlers.Health).Methods(http.MethodGet)
	r.HandleFunc("/users/", handlers.PostUser).Methods(http.MethodPost)
	r.HandleFunc("/users/", handlers.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{email}", handlers.GetUserByEmail).Methods(http.MethodGet)
	r.HandleFunc("/users/", handlers.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{email}", handlers.DeleteUserByEmail).Methods(http.MethodDelete)
}