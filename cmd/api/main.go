package main

import (
	"log"
	"net/http"

	"github.com/aleksanderpalamar/pixpay/config"

	"github.com/aleksanderpalamar/pixpay/internal/router"
	"github.com/aleksanderpalamar/pixpay/pkg/database"
	"github.com/aleksanderpalamar/pixpay/pkg/logger"
	_ "github.com/lib/pq"
)

func main() {
	config.Load()
	logger.InitLogger()

	db, err := database.InitPostgres()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	router := router.NewRouter(db)
	http.Handle("/", router)

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
