package models

import (
	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	Name    string
	Email   string	`json:"email" binding:"required"`
	Phone   string	`json:"phone" binding:"required"`
	Address string
}

func (b *Profile) TableName() string {
	return "Profile"
}