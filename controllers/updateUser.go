package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamkabilan/CRUD-using-Go-and-MYSQL/database"
)

func UpdateUser(responseWriter http.ResponseWriter, request *http.Request) {
	var variables = mux.Vars(request)
	var userId = variables["id"]

	var updateUser User
	parsingErr := json.NewDecoder(request.Body).Decode(&updateUser)
	if parsingErr != nil {
		http.Error(responseWriter, "Invalid Request", http.StatusBadRequest)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	db := database.GetDB()

	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	result, queryErr := db.Exec(query, updateUser.Name, updateUser.Email, userId)
	if queryErr != nil {
		fmt.Println("ERROR (QUERY ERROR): ", queryErr.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	var response Response
	if rowsAffected == 0 {
		response = Response{
			Message: "User with the id " + userId + " does not exist",
		}
		responseWriter.WriteHeader(http.StatusNotFound)
	} else {
		response = Response{
			Message: "User with the id " + userId + " is updated",
		}
	}

	json.NewEncoder(responseWriter).Encode(response)
}
