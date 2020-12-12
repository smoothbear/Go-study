package handler

import (
	"encoding/json"
	"net/http"
	"restful/payload/response"
)

func BookHandler(w http.ResponseWriter, r *http.Request) {
	var book response.Book
	book.Author = "김정빈"
	book.ID = "edy2isoo-asjasnvs"
	book.Name = "Effective Java 3rd Edition"
	book.Title = "[BLACK FRIDAY] Effective Java 3rd Edition"

	res, _ := json.Marshal(book)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}