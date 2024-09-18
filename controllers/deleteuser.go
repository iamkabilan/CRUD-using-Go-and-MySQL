package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamkabilan/CRUD-using-Go-and-MYSQL/database"
)

func DeleteUser(responseWriter http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	userId := variables["id"]

	responseWriter.Header().Set("Content-Type", "application/json")

	db := database.GetDB()

	query := "DELETE FROM users WHERE id = ?"
	result, queryErr := db.Exec(query, userId)

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
			Message: "User with the id " + userId + " is deleted",
		}
	}
	json.NewEncoder(responseWriter).Encode(response)
}
