package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/iamkabilan/CRUD-using-Go-and-MYSQL/database"
)

func findUser(users []User, id int) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}

	return nil
}

func GetUser(responseWriter http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	userId := variables["id"]

	responseWriter.Header().Set("Content-Type", "application/json")

	db := database.GetDB()
	redisConn := database.GetRedis()
	ctx := context.Background()

	var users []User
	var redisUsers []UserRedis
	result := database.GetUsers(redisConn, "users", ctx)

	for _, userJson := range result {
		var user User
		var redisUser UserRedis
		json.Unmarshal([]byte(userJson), &user)
		json.Unmarshal([]byte(userJson), &redisUser)
		users = append(users, user)
		redisUsers = append(redisUsers, redisUser)
	}

	id, _ := strconv.Atoi(userId)
	userExist := findUser(users, int(id))

	if userExist != nil {
		log.Println("User exist in Redis....")
		json.NewEncoder(responseWriter).Encode(userExist)

		var redisUser UserRedis = UserRedis{
			User: User{
				Id:    userExist.Id,
				Name:  userExist.Name,
				Email: userExist.Email,
			},
			Timestamp: time.Now().Unix(),
		}

		isDeleted := database.DeleteKey(redisConn, "users", ctx)
		if isDeleted {
			for _, user := range redisUsers {
				if user.Id != userExist.Id {
					database.InsertUser(redisConn, "users", user, ctx)
				}
			}
		}

		database.InsertUser(redisConn, "user", redisUser, ctx)
	} else {
		log.Println("User does not exist in Redis....")

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
}
