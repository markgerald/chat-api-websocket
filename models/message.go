package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Content  string `json:"content"`
}
