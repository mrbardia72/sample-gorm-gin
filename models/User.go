package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Email   string	`json:"email" binding:"required"`
	Posts []Post `gorm:"ForeignKey:UserID"`
}

func (b *User) TableName() string {
	return "users"
}