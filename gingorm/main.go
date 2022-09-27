package main

import (
	"fmt"
	"gingorm/Config"
	"gingorm/Models"
	"gingorm/Routes"

	"gorm.io/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
