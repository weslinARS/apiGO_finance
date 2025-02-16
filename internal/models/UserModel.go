package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string         `json:"name" validate:"required,min=2,max=50" jsonapi:"attr,name"`
	LastName       string         `json:"lastName" validate:"required,min=2,max=50" jsonapi:"attr,lastName"`
	ID             string         `gorm:"type:uuid;default:uuid_generate_v4()" jsonapi:"attr,id"`
	Email          string         `json:"email" validate:"required,email" jsonapi:"attr,email"`
	Gender         string         `json:"gender" validate:"required,oneof=Female Male" jsonapi:"attr,gender"`
	DateOfBirth    string         `json:"dateOfBirth" gorm:"type:date" validate:"required,datetimeF=2006-01-02" jsonapi:"attr,dateOfBirth" `
	UserCredential UserCredential `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;" json:"userCredential,omitempty" jsonapi:"relation,userCredential"`
	Accounts       []*Account     `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;" json:"accounts,omitempty" jsonapi:"relation,accounts,omitempty"`
	Currencies     []*Currency    `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;" json:"currencies,omitempty" jsonapi:"relation,currencies,omitempty"`
	Categories     []*Category    `gorm:"many2many:user_categories;" json:"categories,omitempty" jsonapi:"relation,categories,omitempty"`
}

func (u *User) JSONAPIMeta() *map[string]interface{} {
	meta := map[string]interface{}{
		"created_at": u.CreatedAt,
	}
	return &meta
}
