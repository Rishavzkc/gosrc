package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*func getbooks(c *gin.Context) {
	var book Book
	if err := c.BindJSON(&book); err != nil {
		return
	}
	Database.Create(&book)
	c.IndentedJSON(http.StatusCreated, book)

} */

//CreateUser ... Create User
/*func CreateBook(c *gin.Context) {
	var book Book
	c.BindJSON(&book)
	err := Database.Create(&book)
	if err != nil {

		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}*/
func getbooks(c *gin.Context) {
	var book []Book
	Database.Find(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
