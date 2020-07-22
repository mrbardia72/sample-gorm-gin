package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title  string
	Body string
	UserID int `gorm:"column:user_id"`
	User   User
}

func (b *Post) TableName() string {
	return "posts"
}