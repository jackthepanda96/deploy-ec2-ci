package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string `json:"title" form:"title"`
	// UserID int    `gorm:"primaryKey"`
}

type BookDBModel struct {
	db *gorm.DB
}

func NewBookModel(db *gorm.DB) *BookDBModel {
	return &BookDBModel{db: db}
}

type BookModel interface {
	Insert(newBook Book) (Book, error)
}

func (u *BookDBModel) Insert(newBook Book) (Book, error) {
	if err := u.db.Save(&newBook).Error; err != nil {
		return newBook, err
	}
	return newBook, nil
}
