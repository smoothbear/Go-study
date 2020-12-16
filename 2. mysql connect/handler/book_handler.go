package handler

import (
	"encoding/json"
	"net/http"
	"restful/mysql/db"
	"restful/mysql/model"
	"restful/mysql/payload/request"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var bookRequest request.Book
	err := decoder.Decode(&bookRequest)
	if err != nil {
		panic(err)
	}

	book := model.Book{Title: bookRequest.Title, Description: bookRequest.Description}
	result := D

}

func DeleteBook(w http.ResponseWriter, r *http.Request){

}