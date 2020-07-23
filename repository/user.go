package repository

import (
	// "encoding/json"
	"fmt"
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mrbardia72/blog-gorm-gorilla/src/helpers"
)
 
//GetAllUsers Fetch all user data
// func GetAllUsers(user *[]models.User) (err error) {
// 	if err = config.DB.Find(user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
 
//CreateProfile ... Insert New data
func CreateUser(user *models.User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetSearchUser(user *models.User, name string) (err error) {

	if err = config.DB.Where("name = ?", name).Find(&user).Error; err != nil {
		info:="There is no such record"
		helpers.LogApi(info)
		return 
	} else { 
		config.DB.Preload("Posts").Find(&user)
		info:=" Successfully search User "
		helpers.LogApi(info)
		return
	}
}

func GetUserByID(user *models.User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
	 return err
	}
	return nil
   }

//UpdateProfile ... Update profile
func UpdateUser(user *models.User, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

//Deleteprofile ... Delete profile
func DeleteUser(user *models.User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
