package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string         `json:"name" validate:"required,min=2,max=50"`
	LastName       string         `json:"lastName" validate:"required,min=2,max=50"`
	ID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email          string         `json:"email" validate:"required,email"`
	Gender         string         `json:"gender" validate:"required,oneof=Female Male"`
	DateOfBirth    string         `json:"dateOfBirth" gorm:"type:date" validate:"required,datetimeF=2006-01-02" `
	UserCredential UserCredential `gorm:"foreignKey:UserId"`
}
