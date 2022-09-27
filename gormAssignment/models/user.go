package models

import (
	"gormAssignment/config"
	"gormAssignment/entities"
)

type UserModel struct {
}

func (userModel UserModel) FindAll() ([]entities.User, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var users []entities.User
		db.Preload("Languages").Find(&users)
		return users, nil
	}
}
