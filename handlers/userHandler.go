package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/D-Toshchakov/pet/social-media/handlers/dto"
	"github.com/D-Toshchakov/pet/social-media/internal/database"
	"github.com/gorilla/mux"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dto.PostUserDto

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can not parse user from body"))
		return
	}

	dbUser := &database.User{
		Email:    user.Email,
		Password: user.Password,
		Age:      user.Age,
		Name:     user.Name,
	}

	dbReq := database.DB.Create(dbUser)
	if dbReq.Error != nil {
		fmt.Println(dbReq.Error)
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("user with this email already exists"))
		return
	}

	// fmt.Printf("New user %#v\n", dbReq)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dbUser)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []database.User

	dbReq := database.DB.Find(&users)
	if dbReq.Error != nil {
		fmt.Println(dbReq.Error)
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Users not found"))
		return
	}

	fmt.Println("all users: ", users)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)

}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	fmt.Println("email to find:", vars["email"])
	var user database.User

	dbReq := database.DB.Where("email = ?", vars["email"]).First(&user)
	if dbReq.Error != nil {
		fmt.Println(dbReq.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user with this email does not exist"))
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser dto.UpdateUserDto

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can not parse user from body"))
		return
	}

	fmt.Println("User to update: ", newUser)

	user := database.User{}

	dbReq := database.DB.Where("email = ?", newUser.Email).First(&user)
	if dbReq.Error != nil {
		fmt.Println(dbReq.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user with this email does not exist"))
		return
	}

	database.DB.Model(&user).Updates(database.User{
		Name:     newUser.Name,
		Password: newUser.Password,
		Age:      newUser.Age,
	})

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)

}

func DeleteUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	fmt.Println("User to delete: ", vars["email"])

	database.DB.Where("email = ?", vars["email"]).Delete(&database.User{})

	w.WriteHeader(http.StatusOK)

	w.Write([]byte{})
}
