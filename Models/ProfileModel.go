package Models

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	Name    string 
	Email   string 
	Phone   string 
	Address string 
}

func (b *Profile) TableName() string {
	return "Profile"
}
