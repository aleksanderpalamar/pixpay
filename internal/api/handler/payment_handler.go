package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type PaymentHandler struct {
	db *sql.DB
}

type Payment struct {
	ID        int64   `json:"id"`
	CreatedAt string  `json:"created_at"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	Pix       string  `json:"pix"`
}

func NewPaymentHandler(db *sql.DB) *PaymentHandler {
	return &PaymentHandler{db: db}
}

func (h *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	err := json.NewDecoder(r.Body).Decode(&payment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	payment.ID = 1
	payment.CreatedAt = time.Now().Format(time.RFC3339)

	_, err = h.db.Exec("INSERT INTO payments (id, created_at, amount, status, pix) VALUES ($1, $2, $3, $4, $5)", payment.ID, payment.CreatedAt, payment.Amount, payment.Status, payment.Pix, payment.Pix)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *PaymentHandler) GetPayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var payment Payment
	err := h.db.QueryRow("SELECT id, created_at, amount, status, pix FROM payments WHERE id = $1", id).Scan(&payment.ID, &payment.CreatedAt, &payment.Amount, &payment.Status, &payment.Pix)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}
