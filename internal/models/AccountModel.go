package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID                 string   `gorm:"type:uuid;default:uuid_generate_v4()" jsonapi:"primary,account"`
	IncludedInBalance  bool     `gorm:"default:true" json:"includedInBalance" jsonapi:"attr,includedInBalance"`
	AccountCutOffDay   int      `gorm:"default:1;check: account_cut_off_day >= 1 AND account_cut_off_day <= 31;" validate:"required,lte:31,gte:" json:"accountCutOffDate" jsonapi:"attr,accountCutOffDate"`
	AccountDeadLineDay int      `gorm:"default:1;check: account_dead_line_day >= 1 AND account_dead_line_day <= 31;" validate:"required,lte:31,gte:" json:"accountDeadlineDay" jsonapi:"attr,accountDeadlineDay"`
	Name               string   `gorm:"not null;size:50" json:"name" validate:"required,min=2,max=50" jsonapi:"attr,name"`
	Balance            float64  `gorm:"not null;default:0" json:"balance" validate:"required" jsonapi:"attr,balance"`
	Type               string   `gorm:"check:type IN ('credit', 'savings');not null" json:"type" validate:"required,oneof=checking savings" jsonapi:"attr,type"`
	CurrencyCode       string   `gorm:"type:char(3);constraint:OnDelete:CASCADE'" json:"currencyCode" validate:"required,len=3" jsonapi:"attr,currencyCode"`
	Currency           Currency `gorm:"foreignkey:CurrencyCode;references:code" json:"currency" validate:"required" jsonapi:"attr,currency"`
	UserId             string   `gorm:"type:uuid;constraint:OnDelete:CASCADE'" json:"userId" validate:"required" jsonapi:"attr,userId"`
	Limit              float64  `gorm:"gte:0" json:"limit" jsonapi:"attr,limit"`
}
