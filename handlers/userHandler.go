package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/D-Toshchakov/pet/social-media/handlers/dto"
	"github.com/D-Toshchakov/pet/social-media/internal/database"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dto.PostUserDto

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can not parse user from body"))
	}
	fmt.Printf("%#v\n", user)

	dbUser := &database.User{
		Email:    user.Email,
		Password: user.Password,
		Age:      user.Age,
		Name:     user.Name,
	}

	pk := database.DB.Create(dbUser)

	fmt.Printf("New user %#v\n", pk)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dbUser)

}

