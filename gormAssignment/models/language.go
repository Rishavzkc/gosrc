package models

import (
	"gormAssignment/config"
	"gormAssignment/entities"
)

type LanguageModel struct {
}

func (languageModel LanguageModel) FindAll() ([]entities.Language, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var languages []entities.Language
		db.Preload("Users").Find(&languages)
		return languages, nil
	}
}
