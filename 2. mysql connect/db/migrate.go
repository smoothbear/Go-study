package db

import (
	"github.com/jinzhu/gorm"
	"restful/mysql/model"
)

func Migrate(db *gorm.DB) {
	db.LogMode(true)

	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
	}

	if !db.HasTable(&model.Book{}) {
		db.CreateTable(&model.Book{})
	}
}
