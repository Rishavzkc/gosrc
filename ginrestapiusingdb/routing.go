package main

import "github.com/gin-gonic/gin"

func HandleRouting() {

	router := gin.Default()

	router.GET("/books", getbooks)

	router.Run()
}
