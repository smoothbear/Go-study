package handler

import (
	"encoding/json"
	"net/http"
	"restful/mysql/payload/request"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var bookRequest request.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bookRequest); err != nil {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(""))
	}
}