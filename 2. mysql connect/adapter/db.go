package adapter

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

func ConnectToMysql() (db *gorm.DB, err error) {
	password := os.Getenv("DB_MYSQL_PASSWORD")
	if password == "" {
		err = errors.New("please set DB_MYSQL_PASSWORD")
		return
	}

	args := fmt.Sprintf("root:%s@(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local", password)
	db, err = gorm.Open("mysql", args)
	return
}
