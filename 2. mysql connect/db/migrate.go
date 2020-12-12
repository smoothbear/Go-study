package db

import "github.com/jinzhu/gorm"

func Migrate(db *gorm.DB) {
	db.LogMode(true)

	if !db.HasTable()
}
