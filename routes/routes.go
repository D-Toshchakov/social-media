package routes

import (
	"github.com/D-Toshchakov/pet/social-media/handlers"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {
	r.HandleFunc("/health", handlers.TestHandler)
}