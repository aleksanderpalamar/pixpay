package database

import (
	"database/sql"
	"fmt"
	"os"
)

func InitPostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", getPostgresConnectionString())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func getPostgresConnectionString() string {
	host := os.Getenv("DB_HOST")
	post := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, post, user, password, dbname)
}
