package models

import "time"

// User represents a user account in the system
type UserModel2 struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // omit from JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate hook to set created/updated time
func (u *UserModel2) BeforeCreate() error {
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = u.CreatedAt
	return nil
}

// BeforeUpdate hook to set updated time
func (u *UserModel2) BeforeUpdate() error {
	u.UpdatedAt = time.Now().UTC()
	return nil 
}