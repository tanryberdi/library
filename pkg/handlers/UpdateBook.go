package handlers

import (
	"encoding/json"
	"io/ioutil"
	"library/pkg/mocks"
	"library/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parametr
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//Read requested body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	//iterate over all the mock books
	for index, book := range mocks.Books {
		if book.Id == id {
			// Update and send a response when book Id matches dynamic id
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author
			book.Desc = updatedBook.Desc

			mocks.Books[index] = book

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
