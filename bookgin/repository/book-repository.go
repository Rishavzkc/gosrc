package repository

import (
	"bookgin/entity"
	"fmt"

	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(b entity.Book) error
	UpdateBook(b entity.Book) error
	DeleteBook(b entity.Book) error
	AllBook() []entity.Book
	FindBookByID(bookID string) entity.Book
}

type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConn,
	}
}

func (db *bookConnection) InsertBook(b entity.Book) error {
	if tx := db.connection.Create(b); tx.Error != nil {
		return fmt.Errorf("failed to insert record in company db: %w", tx.Error)
	}
	return nil

}

func (db *bookConnection) UpdateBook(b entity.Book) error {

	tx := db.connection.Model(entity.Book{ID: b.ID}).Updates(&entity.Book{
		Title:       b.Title,
		Description: b.Description,
	})
	if tx.Error != nil {
		return fmt.Errorf("failed to update record with id '%s' in company db: %w", b.ID, tx.Error)
	}
	return nil
}

func (db *bookConnection) DeleteBook(b entity.Book) error {
	tx := db.connection.Delete(&entity.Book{ID: b.ID})
	if tx.Error != nil {
		return fmt.Errorf("failed to delete record with id '%s': %w", b.ID, tx.Error)
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("no record found to delete for id '%s'", b.ID)
	}
	return nil
}

func (db *bookConnection) FindBookByID(bookID string) entity.Book {
	var book entity.Book
	db.connection.First(&bookID, "id = ?", bookID)

	return book

}

func (db *bookConnection) AllBook() []entity.Book {

	var books []entity.Book
	db.connection.Preload("Book").Find(&books)
	return books
}
