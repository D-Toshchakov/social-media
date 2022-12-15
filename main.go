package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/D-Toshchakov/pet/social-media/internal/database"
	"github.com/gorilla/mux"
)

const address = "localhost:8080"

func main() {
	// Conecting to database
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(database.User{}, database.Post{})
	if err !=nil {
		panic(err)
	}

	fmt.Println("Social Media backend")

	r := mux.NewRouter()

	r.HandleFunc("/", testHandler)
	
	srv := http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(200)
	w.Write([]byte("Hello world"))
}
