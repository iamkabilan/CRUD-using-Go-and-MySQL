package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/iamkabilan/CRUD-using-Go-and-MYSQL/database"
)

func CreateUser(responseWriter http.ResponseWriter, request *http.Request) {
	var MAX_USERS int64 = 5
	responseWriter.Header().Set("Content-Type", "application/json")

	var createUser User
	parsingErr := json.NewDecoder(request.Body).Decode(&createUser)
	if parsingErr != nil {
		http.Error(responseWriter, "Invalid Request", http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	redisConn := database.GetRedis()
	ctx := context.Background()

	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, queryErr := db.Exec(query, createUser.Name, createUser.Email)
	if queryErr != nil {
		fmt.Println("ERROR (QUERY ERROR): ", queryErr.Error())
		return
	}

	id, _ := result.LastInsertId()
	createUser.Id = int(id)

	var userRedis UserRedis
	userRedis = UserRedis{
		User:      createUser,
		Timestamp: time.Now().Unix(),
	}
	userJson, _ := json.Marshal(userRedis)
	key := "users"

	size := database.KeySize(redisConn, key, ctx)
	if size != -1 && size >= MAX_USERS {
		database.RemoveUser(redisConn, key, ctx)
		database.InsertUser(redisConn, key, userJson, ctx)
	} else if size != -1 && size < MAX_USERS {
		database.InsertUser(redisConn, key, userJson, ctx)
	} else {
		fmt.Println("Redis Error")
	}

	json.NewEncoder(responseWriter).Encode(createUser)
}
