package main

import (
	"fmt"
	"log"
	"userdata/factory"
	"userdata/models"
	"userdata/utilities"
)

func main() {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}

	userDao := factory.FactoryDao(config.Engine)

	user := models.User{}
	fmt.Println("Enter first name:")
	fmt.Scan(&user.FirstName)
	fmt.Println("Enter last name:")
	fmt.Scan(&user.LastName)
	fmt.Println("Enter Email:")
	fmt.Scan(&user.Email)

	err = userDao.Create(&user)

	if err != nil {
		log.Fatal(err)
		return
	}

}
