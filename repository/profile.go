package repository

import (
	"fmt"
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllProfiles(profile *[]models.Profile) (err error) {
	if err = config.DB.Find(profile).Error; err != nil {
		return err
	}
	return nil
}

//CreateProfile ... Insert New data
func CreateProfile(profile *models.Profile) (err error) {
	if err = config.DB.Create(profile).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetSearch(profile *models.Profile, name string) (err error) {
	if err = config.DB.Where("name = ?", name).First(profile).Error; err != nil {
		return err
	}
	return nil
}

func GetProfileByID(profile *models.Profile, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(profile).Error; err != nil {
	 return err
	}
	return nil
   }

//UpdateProfile ... Update profile
func UpdateProfile(profile *models.Profile, id string) (err error) {
	fmt.Println(profile)
	config.DB.Save(profile)
	return nil
}

//Deleteprofile ... Delete profile
func DeleteProfile(profile *models.Profile, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(profile)
	return nil
}
