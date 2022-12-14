package main

import (
	"gormexample/database"
	"gormexample/routes"
)

func main() {

	database.InitDb()

	defer database.CloseDbConnextion(database.ConnectDB())

	r := routes.SetupRouter()
	_ = r.Run(":8080")
}
