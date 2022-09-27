package Models

import (
	"fmt"
	"gingorm/Config"
)

// Get all book data
func GetAllBooks(book *[]Book) (err error) {
	if err = Config.DB.Find(book).Error; err != nil {
		return err
	}
	return nil
}

//Insert New data
func CreateBook(book *Book) (err error) {
	if err = Config.DB.Create(book).Error; err != nil {
		return err
	}
	return nil
}

//get only one book by Id
func GetBookByID(book *Book, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(book).Error; err != nil {
		return err
	}
	return nil
}

//Update book
func UpdateBook(book *Book, id string) (err error) {
	fmt.Println(book)
	Config.DB.Save(book)
	return nil
}

// Delete book
func DeleteBook(book *Book, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(book)
	return nil
}
