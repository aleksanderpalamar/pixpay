package pix

import (
	"encoding/json"
	"net/http"
)

type PixHandler struct {
	Service PixService
}

func NewPixHandler(service PixService) *PixHandler {
	return &PixHandler{Service: service}
}

func (h *PixHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment PixPayment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.CreatePayment(payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
