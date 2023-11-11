package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
  	Email string `json:"email" gorm:"unique;not null"`
	Password  string  `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate hook to set created/updated time
func (u *User) BeforeCreate() error {
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = u.CreatedAt
	return nil
}

// BeforeUpdate hook to set updated time
func (u *User) BeforeUpdate() error {
	u.UpdatedAt = time.Now().UTC()
	return nil 
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) 
  }