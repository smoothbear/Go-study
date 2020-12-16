package db

import (
	"gorm.io/gorm"
	"restful/mysql/model"
)

type Accessor interface {
	// Create Method
	CreateUser(user *model.User) (resultUser *model.User)
	CreateBook(book *model.Book)

	// Get Method
	GetUser(userId string)
	GetBook(UUID string)

	// Delete Method
	DeleteUser(userId string) error
	DeleteBook(UUID string) error

	// DB Transaction
	BeginTx()
	Commit() *gorm.DB
	RollBack() *gorm.DB
}