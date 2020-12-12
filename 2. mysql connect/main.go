package main

import (
	"github.com/gorilla/mux"
	"restful/mysql/handler"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/book", handler.CreateBook).Methods("POST")
}

