package main

import (
	"fmt"
	"os"
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load();
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println("CRUD Application");

	var host string = os.Getenv("MYSQL_HOST");
	var username string = os.Getenv("MYSQL_USERNAME");
	var password string = os.Getenv("MYSQL_PASSWORD");
	var port string = os.Getenv("MYSQL_PORT");


	var dsn string = username+":"+password+"@tcp("+host+":"+port+")/users";
	fmt.Println(dsn);
	db, db_err := sql.Open("mysql", dsn);
	if db_err != nil {
		fmt.Println("ERROR: ",err);
	}
	defer db.Close();


}
