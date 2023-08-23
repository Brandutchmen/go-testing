package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Username string    `gorm:"uniqueIndex;not null;size:50;" validate:"required,min=3,max=50" json:"username"`
	Email    string    `gorm:"uniqueIndex;not null;size:255;" validate:"required,email" json:"email"`
	Password string    `gorm:"not null;" validate:"required,min=6,max=50" json:"password"`
	Names    string    `json:"names"`
}
