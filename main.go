package main

import (
	"github.com/mrbardia72/sample-gorm-gin/Config"
	"github.com/mrbardia72/sample-gorm-gin/Models"
	"github.com/mrbardia72/sample-gorm-gin/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Profile{})

	r := Routes.SetupRouter()
	//running
	r.Run()
}
