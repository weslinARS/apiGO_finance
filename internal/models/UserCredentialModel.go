package models

import (
	"gorm.io/gorm"
)

type UserCredential struct {
	gorm.Model
	UserEmail string `json:"userEmail" jsonapi:"attr,userEmail"`
	Password  string `json:"password" jsonapi:"attr,password"`
	UserId    string `gorm:"type:uuid" json:"userId" jsonapi:"attr,userId"`
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();constraint:OnDelete:CASCADE;" json:"id" jsonapi:"primary,userCredentials"`
}
