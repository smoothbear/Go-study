package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UUID uuid `gorm:"PRIMARY_KEY;Type:char(20);UNIQUE;INDEX" validate:"uuid=user, len=15"`		// form => 'user-' + 11 digits number (15 length)
	UserId string `gorm:"Type:varchar(20);UNIQUE;NOT NULL" validate:"min=4,max=20,ascii"`		// between 4 ~ 15
	Password string `gorm:"Type:varchar(25);NOT NULL" validate:"min=4,max=25,ascii"`			// between 4 ~ 20
}

type Book struct {
	gorm.Model
	UUID uuid `gorm:"PRIMARY_KEY;Type:char(25);UNIQUE;INDEX" validate:"uuid=book, len=20"` // form => 'book-' + 16 digits number (20 length)
	Title string `gorm:"Type:varchar(20);NOT NULL" validate:"max=20,ascii"`
	Description string `gorm:"varchar(20);NOT NULL" validate:"max=20,ascii"`
}