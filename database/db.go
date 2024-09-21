package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var (
	db      *sql.DB
	redisDb *redis.Client
)

func Initialize() error {
	var err error
	db, err = ConnectToDatabase()
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	log.Println("Connected to the MySQL Database.")

	redisDb = connectToRedis()
	pong, err := redisDb.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to the Redis: %w", err)
	}
	log.Println("Connected to the Redis, ", pong)

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

func connectToRedis() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	redisDb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "",
		DB:       0,
	})

	return redisDb
}
