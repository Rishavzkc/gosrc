package config

import (
	"gormAssignment/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var urlDSN = "root:Quest1234@tcp(localhost:3306)/gorm?parseTime=true"

func GetDB() (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities.User{}, &entities.Language{})
	return db, nil

}
