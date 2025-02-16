package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model
	UserId   string `gorm:"type:uuid;constraint:OnDelete:CASCADE'" json:"userId" validate:"required" jsonapi:"attr,userId"`
	Name     string `gorm:"not null;unique;index" json:"name" validate:"required" jsonapi:"attr,name"`
	Code     string `gorm:"not null;unique;size:3" json:"code" validate:"required,len=3" jsonapi:"attr,code"`
	Symbol   string `gorm:"not null;size:5" json:"symbol" validate:"required" jsonapi:"attr,symbol"`
	Decimals int    `gorm:"not null;default:2" json:"decimals" validate:"required,min=0,max=4" jsonapi:"attr,decimals"`
}
