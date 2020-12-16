package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"restful/mysql/adapter"
	"restful/mysql/db"
	"restful/mysql/handler"
)

var DB *gorm.DB

// Main Server Start
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/book", handler.CreateBook).Methods("POST")

	var err error

	dbc, err := adapter.ConnectToMysql()

	if err != nil {
		fmt.Print("Failed to connect database")
		panic(err)
	}

	db.Migrate(dbc)

	defer DB.Close()

	log.Fatal(http.ListenAndServe(":8000", router))
}
