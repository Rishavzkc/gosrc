package main

import (
	"fmt"
	"gormAssignment/models"
)

// func main() {
// 	fmt.Println("List the languages for each user")
// 	ListLanguagesForEachUser()

// 	fmt.Println("List the users for each language")
// 	ListUsersForEachLanguage()
// }

func ListLanguagesForEachUser() {
	var userModel models.UserModel

	users, _ := userModel.FindAll()
	for _, user := range users {
		fmt.Println(user.ToString())
		fmt.Println("Languages: ", len(user.Languages))
		if len(user.Languages) > 0 {
			for _, language := range user.Languages {
				fmt.Println(language.ToString())
				fmt.Println("========================")
			}
		}
		fmt.Println("-----------------------------")
	}
}

func ListUsersForEachLanguage() {
	var languageModel models.LanguageModel
	languages, _ := languageModel.FindAll()
	for _, language := range languages {
		fmt.Println(language.ToString())
		fmt.Println("Users: ", len(language.Users))
		if len(language.Users) > 0 {
			for _, user := range language.Users {
				fmt.Println(user.ToString())
				fmt.Println("========================")
			}
		}
		fmt.Println("-----------------------------")
	}
}
