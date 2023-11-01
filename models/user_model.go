package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"username" validate:"required"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password"`
}

// Password  string    `json:"-"` // omit from JSON
// Password  string  `json:"password" binding:"required,min=6"`
// CreatedAt time.Time `json:"created_at"`
// UpdatedAt time.Time `json:"updated_at"`
// BeforeCreate hook to set created/updated time
// func (u *User) BeforeCreate() error {
// 	u.CreatedAt = time.Now().UTC()
// 	u.UpdatedAt = u.CreatedAt
// 	return nil
// }

// BeforeUpdate hook to set updated time
// func (u *User) BeforeUpdate() error {
// 	u.UpdatedAt = time.Now().UTC()
// 	return nil 
// }

// func NewUser(userName, firstName, lastName, email, password string) *User {
	// Generate a new UUID for the user's ID
	// id := uuid.New()

	// Create a new User instance
// 	user := &User{
// 		ID:        id,
// 		UserName:  userName,
// 		FirstName: firstName,
// 		LastName:  lastName,
// 		Email:     email,
// 		Password:  password,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	return user
// }