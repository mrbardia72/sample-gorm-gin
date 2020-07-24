package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title  string
	Body string
	UserID int `gorm:"column:user_id"`
	Users User `gorm:"ForeignKey:UserID"`
}

func (b *Post) TableName() string {
	return "posts"
}