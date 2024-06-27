package main

import (
	"log"
	"net/http"

	"github.com/aleksanderpalamar/pixpay/internal/database"
	"github.com/aleksanderpalamar/pixpay/internal/pix"

	_ "github.com/lib/pq"
)

func main() {
	db := database.Connect()
	defer db.Close()

	pixService := pix.NewPixService(db)
	pixHandler := pix.NewPixHandler(pixService)

	http.HandleFunc("/payments", pixHandler.CreatePayment)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
