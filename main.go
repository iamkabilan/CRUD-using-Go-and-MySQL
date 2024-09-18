package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/iamkabilan/CRUD-using-Go-and-MYSQL/controllers"
	"github.com/iamkabilan/CRUD-using-Go-and-MYSQL/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	if err := database.Initialize(); err != nil {
		fmt.Println(err)
		return
	}
	defer database.GetDB().Close()

	fmt.Println("CRUD Application")

	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/createuser", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/deleteuser/{id}", controllers.DeleteUser).Methods("DELETE")

	var PORT = os.Getenv("PORT")
	log.Printf("Starting server on port %s", PORT)
	http.ListenAndServe(":"+PORT, router)
}
