package router

import (
	"database/sql"

	"github.com/aleksanderpalamar/pixpay/internal/api/handler"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	h := handler.NewPaymentHandler(db)

	router.HandleFunc("/payments", h.CreatePayment).Methods("POST")
	router.HandleFunc("/payments/{id}", h.GetPayment).Methods("GET")

	return router
}
