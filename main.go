package main

import (
	"fmt"
	"os"
	"database/sql"
	"net/http"
	"log"
	"encoding/json"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	Id int	`json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func connectToDatabase() (*sql.DB, error) {
	var host string = os.Getenv("MYSQL_HOST");
	var username string = os.Getenv("MYSQL_USERNAME");
	var password string = os.Getenv("MYSQL_PASSWORD");
	var port string = os.Getenv("MYSQL_PORT");


	var dsn string = username+":"+password+"@tcp("+host+":"+port+")/CRUD";
	db, db_err := sql.Open("mysql", dsn);
	if db_err != nil {
		fmt.Println("ERROR: ",db_err);
	}

	return db, nil;
}

func getUsers(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json");

	db, dbErr := connectToDatabase();
	if dbErr != nil {
		fmt.Println("ERROR (DB ERROR): ",dbErr);
	}
	defer db.Close();

	rows, queryErr := db.Query("SELECT id, name, email FROM users");
	if queryErr != nil {
		fmt.Println("ERROR (QUERY ERROR): ",queryErr);
		return;
	}
	defer rows.Close();
	
	var users []User;

	for rows.Next() {
		var user User;
		scanErr := rows.Scan(&user.Id, &user.Name, &user.Email);
		if scanErr != nil {
			fmt.Println("ERROR: ",scanErr);
		}

		users = append(users, user);
	}

	json.NewEncoder(responseWriter).Encode(users);
}

func createUser(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json");

	var createUser User;
	parsingErr := json.NewDecoder(request.Body).Decode(&createUser);
	if parsingErr != nil {
		http.Error(responseWriter, "Invalid Request", http.StatusBadRequest);
		return;
	}

	db, dbErr := connectToDatabase();
	if dbErr != nil {
		fmt.Println("ERROR (DB ERROR): ",dbErr);
		return;
	}
	defer db.Close();

	query := "INSERT INTO users (name, email) VALUES (?, ?)";
	result, queryErr := db.Exec(query, createUser.Name, createUser.Email);
	if queryErr != nil {
		fmt.Println("ERROR (QUERY ERROR): ", queryErr.Error());
	}

	id, _ := result.LastInsertId();
	createUser.Id = int(id);

	json.NewEncoder(responseWriter).Encode(createUser);
}

func main() {
	err := godotenv.Load();
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println("CRUD Application");

	router := mux.NewRouter();
	router.HandleFunc("/users", getUsers).Methods("GET");
	router.HandleFunc("/createuser", createUser).Methods("POST");


	var PORT = os.Getenv("PORT");
	log.Printf("Starting server on port %s", PORT);
	http.ListenAndServe(":"+PORT, router);
}
