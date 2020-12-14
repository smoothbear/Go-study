package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restful/mysql/handler"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/book", handler.CreateBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

