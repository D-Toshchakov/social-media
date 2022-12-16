package routes

import (
	"net/http"

	"github.com/D-Toshchakov/pet/social-media/handlers"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {
	r.HandleFunc("/health", handlers.Health).Methods(http.MethodGet)
	r.HandleFunc("/users", handlers.PostUser).Methods(http.MethodPost)
}