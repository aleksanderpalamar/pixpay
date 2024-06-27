package database

import (
	"database/sql"
	"fmt"
	"log"
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

func runMigrations(db *sql.DB) error {
	file, err := os.ReadFile("scripts/init.sql")
	if err != nil {
		return fmt.Errorf("could not read init.sql file: %v", err)
	}
	_, err = db.Exec(string(file))
	if err != nil {
		return fmt.Errorf("could not execute migrations: %v", err)
	}
	log.Println("Database migrations executed successfully")
	return nil
}
