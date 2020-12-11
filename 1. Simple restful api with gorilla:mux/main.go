package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restful/handler"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.BookHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
