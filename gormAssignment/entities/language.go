package entities

import "fmt"

type Language struct {
	Id    int `gorm:"primary_key, AUTO_INCREMENT"`
	Name  string
	Users []User `gorm:"many2many:user_language"`
}

func (language *Language) TableName() string {
	return "language"
}

func (language Language) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s", language.Id, language.Name)
}