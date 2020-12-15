package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"restful/mysql/handler"

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

	db, err := gorm.Open("gorm.db", &gorm.Confi)
	}

	log.Fatal(http.ListenAndServe(":8000", app.Router))
}
