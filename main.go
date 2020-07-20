package main

import (
	"github.com/mrbardia72/sample-gorm-gin/config"
	"github.com/mrbardia72/sample-gorm-gin/models"
	"github.com/mrbardia72/sample-gorm-gin/routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error
 
func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	
	config.DB.AutoMigrate(&models.Profile{})

	r := routes.SetupRouter() 

	r.Run()
}
