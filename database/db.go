package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Initialize() error {
	var err error
	db, err = ConnectToDatabase()
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	return nil
}

func GetDB() *sql.DB {
	return db
}

func ConnectToDatabase() (*sql.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	port := os.Getenv("MYSQL_PORT")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/CRUD"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("ERROR: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ERROR: %v", err)
	}

	return db, nil
}
