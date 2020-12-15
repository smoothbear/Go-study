package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"restful/mysql/handler"
	"gorm.io/driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	DB gorm.DB
	Router *mux.Router
}

// Main Server Start
func main() {
	var app App
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/book", handler.CreateBook).Methods("POST")

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:@tcp(127.0.0.1:3306)/go-study",
		DefaultStringSize: 256,
	}))
	}

	log.Fatal(http.ListenAndServe(":8000", app.Router))
}
