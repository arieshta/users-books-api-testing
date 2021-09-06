package database

import (
	"users-books-api-testing/config"
	"users-books-api-testing/models"
)

func AddBook(book *models.Books) error {
	if err := config.DB.Table("books").Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func GetBooks() (interface{}, error) {
	var books []models.Books

	if err := config.DB.Table("books").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookById(id int) (interface{}, error) {
	var book models.Books

	if err := config.DB.Table("books").First(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBookById(id int, book *models.Books) error {
	var books models.Books
	if err := config.DB.Table("books").First(&books, id).Error; err != nil {
		return err
	}
	err := config.DB.Table("books").Where("id = ?", id).Updates(book).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteBookById(id int) error {
	var book models.Books
	if err := config.DB.Table("books").Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}
	return nil
}
