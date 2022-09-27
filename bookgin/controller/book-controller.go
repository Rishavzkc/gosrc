package controller

import (
	//"bookgin/dto"

	"bookgin/entity"
	"bookgin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

//NewBookController create a new instances of BoookController
func NewBookController(bookServ service.BookService) BookController {
	return &bookController{
		bookService: bookServ,
	}
}

func (c *bookController) All(context *gin.Context) {
	var books []entity.Book = c.bookService.All()
	context.JSON(http.StatusOK, books)
}

func (c *bookController) FindByID(context *gin.Context) {
	_, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found"})
		return
	}
	var book entity.Book = c.bookService.FindByID("id")
	context.JSON(http.StatusOK, gin.H{"data": book})
}

func (c *bookController) Insert(context *gin.Context) {
	var book entity.Book
	context.ShouldBind(&book)

	c.bookService.Insert(book)
	context.JSON(http.StatusCreated, book)

}

func (c *bookController) Update(context *gin.Context) {
	var bookUpdateDTO entity.Book
	errDTO := context.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {

		context.JSON(http.StatusBadRequest, errDTO)
		return
	}
	result := c.bookService.Update(bookUpdateDTO)

	context.JSON(http.StatusOK, result)

	context.JSON(http.StatusForbidden, bookUpdateDTO)

}

func (c *bookController) Delete(context *gin.Context) {
	var book entity.Book
	_, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {

		context.JSON(http.StatusBadRequest, err.Error())

		book.ID = book.ID
		c.bookService.Delete(book)

		context.JSON(http.StatusOK, gin.H{"message": "Deleted Sucessfully"})
	}

}
