package database

import (
	"gorm.io/gorm"
)

// User -
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Posts    []Post `gorm:"foreignKey:ID;references:ID"`
}

// Post -
type Post struct {
	gorm.Model
	Text     string   `json:"text"`
}
