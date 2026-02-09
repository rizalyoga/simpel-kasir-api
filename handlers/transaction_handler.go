package handlers

import (
	"encoding/json"
	"net/http"

	resp "kasir-api-bootcamp/common/handlers"
	"kasir-api-bootcamp/models"
	"kasir-api-bootcamp/services"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		resp.WriteJSON(w, http.StatusMethodNotAllowed, resp.Response{
			Status:  "error",
			Code:    http.StatusMethodNotAllowed,
			Message: "Method not allowed",
			Data:    nil,
		})
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, resp.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		resp.WriteJSON(w, http.StatusInternalServerError, resp.Response{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	resp.WriteJSON(w, http.StatusCreated, resp.Response{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Checkout successful",
		Data:    transaction,
	})
}
