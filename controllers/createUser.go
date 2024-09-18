package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iamkabilan/CRUD-using-Go-and-MYSQL/database"
)

func CreateUser(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	var createUser User
	parsingErr := json.NewDecoder(request.Body).Decode(&createUser)
	if parsingErr != nil {
		http.Error(responseWriter, "Invalid Request", http.StatusBadRequest)
		return
	}

	db := database.GetDB()

	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, queryErr := db.Exec(query, createUser.Name, createUser.Email)
	if queryErr != nil {
		fmt.Println("ERROR (QUERY ERROR): ", queryErr.Error())
		return
	}

	id, _ := result.LastInsertId()
	createUser.Id = int(id)

	json.NewEncoder(responseWriter).Encode(createUser)
}
