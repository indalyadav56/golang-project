package models

import "time"


type Message struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}