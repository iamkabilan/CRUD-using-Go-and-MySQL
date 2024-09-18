package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iamkabilan/CRUD-using-Go-and-MYSQL/database"
)

func GetUsers(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	db := database.GetDB()

	rows, queryErr := db.Query("SELECT id, name, email FROM users")
	if queryErr != nil {
		fmt.Println("ERROR (QUERY ERROR): ", queryErr)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		scanErr := rows.Scan(&user.Id, &user.Name, &user.Email)
		if scanErr != nil {
			fmt.Println("ERROR: ", scanErr)
		}

		users = append(users, user)
	}

	json.NewEncoder(responseWriter).Encode(users)
}
