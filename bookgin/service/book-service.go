package service

import (
	"bookgin/config"
	"bookgin/entity"
	"bookgin/repository"
	"log"

	"github.com/mashingan/smapping"
)

type BookService interface {
	Insert(b entity.Book) entity.Book
	Update(b entity.Book) error
	Delete(b entity.Book)
	All() ([]entity.Book)
	FindByID(bookID string) entity.Book
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (service *bookService) Insert(b entity.Book) entity.Book {
	book := entity.Book{}

	book = entity.Book{ID: book.ID, Title: book.Title, Description: book.Description}
	config.SetupDatabaseConnection().Create(&book)
	return book

}

func (service *bookService) Update(b entity.Book) error {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bookRepository.UpdateBook(book)
	return res
}

func (service *bookService) Delete(b entity.Book) {
	service.bookRepository.DeleteBook(b)
}

func (service *bookService) All() ([]entity.Book ){
	return service.bookRepository.AllBook()
}

func (service *bookService) FindByID(bookID string) entity.Book {
	return service.bookRepository.FindBookByID(bookID)
}
