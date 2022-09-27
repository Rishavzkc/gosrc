package Controllers

import (
	"fmt"
	"gingorm/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get all books
func GetBooks(c *gin.Context) {
	var book []Models.Book
	err := Models.GetAllBooks(&book)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

// Create Book
func CreateBook(c *gin.Context) {
	var book Models.Book
	c.BindJSON(&book)
	err := Models.CreateBook(&book)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

//Get book by id
func GetBookByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetBookByID(&book, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

// Update book details
func UpdateBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.GetBookByID(&book, id)
	if err != nil {
		c.JSON(http.StatusNotFound, book)
	}
	c.BindJSON(&book)
	err = Models.UpdateBook(&book, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

// Delete  Book
func DeleteBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.DeleteBook(&book, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
