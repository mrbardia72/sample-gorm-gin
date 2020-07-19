package Models

import (
	"fmt"
	"github.com/mrbardia72/sample-gorm-gin/Config"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllProfiles(profile *[]Profile) (err error) {
	if err = Config.DB.Find(profile).Error; err != nil {
		return err
	}
	return nil
}

//CreateProfile ... Insert New data
func CreateProfile(profile *Profile) (err error) {
	if err = Config.DB.Create(profile).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetSearch(profile *Profile, name string) (err error) {
	if err = Config.DB.Where("name = ?", name).First(profile).Error; err != nil {
		return err
	}
	return nil
}

func GetProfileByID(profile *Profile, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(profile).Error; err != nil {
	 return err
	}
	return nil
   }

//UpdateProfile ... Update profile
func UpdateProfile(profile *Profile, id string) (err error) {
	fmt.Println(profile)
	Config.DB.Save(profile)
	return nil
}

//Deleteprofile ... Delete profile
func DeleteProfile(profile *Profile, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(profile)
	return nil
}
