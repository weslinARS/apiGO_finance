package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCredential struct {
	gorm.Model
	UserEmail string    `json:"userEmail"`
	Password  string    `json:"password"`
	UserId    uuid.UUID `json:"userId"`
}
