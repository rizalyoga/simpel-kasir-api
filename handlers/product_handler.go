package handlers

import (
	"encoding/json"
	resp "kasir-api-bootcamp/common/handlers"
	"kasir-api-bootcamp/models"
	"kasir-api-bootcamp/services"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		resp.WriteJSON(w, http.StatusMethodNotAllowed, resp.Response{
			Status:  "error",
			Code:    http.StatusMethodNotAllowed,
			Message: "Method not allowed",
			Data:    nil,
		})
	}
}

func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		resp.WriteJSON(w, http.StatusMethodNotAllowed, resp.Response{
			Status:  "error",
			Code:    http.StatusMethodNotAllowed,
			Message: "Method not allowed",
			Data:    nil,
		})
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	products, err := h.service.GetAll(name)
	if err != nil {
		resp.WriteJSON(w, http.StatusInternalServerError, resp.Response{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, resp.Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Product list",
		Data:    products,
	})
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, resp.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		resp.WriteJSON(w, http.StatusNotFound, resp.Response{
			Status:  "error",
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, resp.Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Product details",
		Data:    product,
	})
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, resp.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	err := h.service.Create(&product)
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
		Message: "Product added successfully",
		Data:    product,
	})
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, resp.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, resp.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	product.ID = id
	err = h.service.Update(&product)
	if err != nil {
		resp.WriteJSON(w, http.StatusNotFound, resp.Response{
			Status:  "error",
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, resp.Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Product updated successfully",
		Data:    product,
	})
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, resp.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		resp.WriteJSON(w, http.StatusNotFound, resp.Response{
			Status:  "error",
			Code:    http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	resp.WriteJSON(w, http.StatusOK, resp.Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Product deleted successfully",
		Data:    nil,
	})
}
