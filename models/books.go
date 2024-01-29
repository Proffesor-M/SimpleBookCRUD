package models

import (
	"SampleCRM/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `form:"title" json:"title"`
	Author string `form:"author" json:"author"`
}

func SearchBooks(searchParams Book) ([]Book, error) {
	var books []Book
	if err := config.DB.Where(searchParams).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func CreateBook(book *Book) error {
	if err := config.DB.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(id string, book *Book) error {
	if err := config.DB.Model(&Book{}).Where("id = ?", id).Updates(book).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBook(id string) error {
	if err := config.DB.Delete(&Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
