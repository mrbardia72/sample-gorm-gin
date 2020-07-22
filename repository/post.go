package repository

import (
	// "fmt"
	"github.com/mrbardia72/sample-gorm-gin/config"
	 "github.com/mrbardia72/sample-gorm-gin/models"
	_ "github.com/go-sql-driver/mysql"
)  

//GetAllposts Fetch all user data
func GetAllposts(post *[]models.Post) (err error) {
	if err = config.DB.Find(post).Error; err != nil {
		return err
	}
	return nil
}

//Createpost ... Insert New data
// func Createpost(post *models.Post) (err error) {
// 	if err = config.DB.Create(post).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

//GetSearchPost ... Fetch only one user by Id
// func GetSearchPost(post *models.Post, title string) (err error) {
// 	if err = config.DB.Where("title = ?", title).First(post).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetpostByID(post *models.Post, id string) (err error) {
// 	if err = config.DB.Where("id = ?", id).First(post).Error; err != nil {
// 	 return err
// 	}
// 	return nil
//    }

//Updatepost ... Update post
// func Updatepost(post *models.Post, id string) (err error) {
// 	fmt.Println(post)
// 	config.DB.Save(post)
// 	return nil
// }

//Deletepost ... Delete post
// func Deletepost(post *models.Post, id string) (err error) {
// 	config.DB.Where("id = ?", id).Delete(post)
// 	return nil
// }
