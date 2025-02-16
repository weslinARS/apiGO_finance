package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID          string `gorm:"type:uuid;default:uuid_generate_v4()" jsonapi:"primary,category"`
	Name        string `gorm:"size:50;not null;" validate:"required,min=3,max=50" json:"name" jsonapi:"attr,name"`
	Description string `gorm:"type:text" validate:"required,min=3,max=50" jsonapi:"attr,description" json:"description"`
	IsDefault   bool   `gorm:"default:false" json:"isDefault" jsonapi:"attr,isDefault"`
	Users       []User `gorm:"many2many:user_categories;" json:"users,omitempty" jsonapi:"relation,users"`
}

/*
type UserCategory struct {
	CategoryId string         `gorm:"type:uuid;constraint:OnDelete:CASCADE;primaryKey" json:"categoryId" validate:"required" jsonapi:"attr,categoryId"`
	UserId     string         `gorm:"type:uuid;constraint:OnDelete:CASCADE;primaryKey" json:"userId" validate:"required" jsonapi:"attr,userId"`
	CreatedAt  time.Time      `json:"createdAt" jsonapi:"attr,createdAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" jsonapi:"attr,deletedAt"`
}
*/

func (c *Category) Meta() *map[string]interface{} {
	meta := map[string]interface{}{
		"created_at": c.CreatedAt,
	}
	return &meta
}
