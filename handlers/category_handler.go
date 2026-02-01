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

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
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

func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
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

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
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
		Message: "Category list",
		Data:    categories,
	})
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
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
	category, err := h.service.GetByID(id)
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
		Message: "Category details",
		Data:    category,
	})
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, resp.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}
	err := h.service.Create(&category)
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
		Message: "Category added successfully",
		Data:    category,
	})
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
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
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		resp.WriteJSON(w, http.StatusBadRequest, resp.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}
	category.ID = id
	err = h.service.Update(&category)
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
		Message: "Category updated successfully",
		Data:    category,
	})
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
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
		Message: "Category deleted successfully",
		Data:    nil,
	})
}
