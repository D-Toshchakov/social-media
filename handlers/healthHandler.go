package handlers

import "net/http"

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(200)
	w.Write([]byte("I'm healthy"))
}