package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamkabilan/CRUD-using-Go-and-MYSQL/database"
)

func GetUser(responseWriter http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	userId := variables["id"]

	responseWriter.Header().Set("Content-Type", "application/json")

	db := database.GetDB()

	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := db.QueryRow(query, userId)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			responseWriter.WriteHeader(http.StatusNotFound)
			json.NewEncoder(responseWriter).Encode(map[string]string{
				"error": "User not found",
			})
			return
		}

		responseWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(responseWriter).Encode(map[string]string{
			"error": "Internal server error",
		})
		return
	}

	json.NewEncoder(responseWriter).Encode(user)
}
