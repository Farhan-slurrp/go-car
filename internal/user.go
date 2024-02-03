package internal

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Role  string `json:"role"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthorizeResponse struct {
	User *User `json:"user",omitempty`
}
