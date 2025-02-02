package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string         `json:"name" validate:"required,min=2,max=50" jsonapi:"attr,name"`
	LastName       string         `json:"lastName" validate:"required,min=2,max=50" jsonapi:"attr,lastName"`
	ID             string         `gorm:"type:uuid;default:uuid_generate_v4()" jsonapi:"primary,users"`
	Email          string         `json:"email" validate:"required,email" jsonapi:"attr,email"`
	Gender         string         `json:"gender" validate:"required,oneof=Female Male" jsonapi:"attr,gender"`
	DateOfBirth    string         `json:"dateOfBirth" gorm:"type:date" validate:"required,datetimeF=2006-01-02" jsonapi:"attr,dateOfBirth" `
	UserCredential UserCredential `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	Accounts       []Account      `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
}

func (u *User) JSONAPIMeta() *map[string]interface{} {
	meta := map[string]interface{}{
		"created_at": u.CreatedAt,
	}
	return &meta
}
