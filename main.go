package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/D-Toshchakov/pet/social-media/internal/database"
	"github.com/D-Toshchakov/pet/social-media/routes"
	"github.com/gorilla/mux"
)

const address = "localhost:8080"

func main() {
	// Conecting to database
	err := database.Connect()
	if err != nil {
		panic(err)
	}

	fmt.Println("Social Media backend")

	r := mux.NewRouter()
	
	routes.Setup(r)

	srv := http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
