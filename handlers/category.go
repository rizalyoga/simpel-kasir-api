package handlers

import (
	"encoding/json"
	"kasir-api-bootcamp/common/errors"
	"kasir-api-bootcamp/common/handlers"
	"kasir-api-bootcamp/models"
	"kasir-api-bootcamp/services"
	"net/http"
	"strconv"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := services.GetCategories()
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, response.Response{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Category list",
		Data:    categories,
	})
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	category, err := services.GetCategory(id)
	if err != nil {
		if _, ok := err.(*errors.ErrNotFound); ok {
			response.WriteJSON(w, http.StatusNotFound, response.Response{
				Status:  "error",
				Code:    http.StatusNotFound,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
		response.WriteJSON(w, http.StatusInternalServerError, response.Response{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Category details",
		Data:    category,
	})
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	category, err := services.CreateCategory(newCategory)
	if err != nil {
		response.WriteJSON(w, http.StatusInternalServerError, response.Response{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	response.WriteJSON(w, http.StatusCreated, response.Response{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "Category added successfully",
		Data:    category,
	})
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	var updatedCategory models.Category
	if err := json.NewDecoder(r.Body).Decode(&updatedCategory); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	category, err := services.UpdateCategory(id, updatedCategory)
	if err != nil {
		if _, ok := err.(*errors.ErrNotFound); ok {
			response.WriteJSON(w, http.StatusNotFound, response.Response{
				Status:  "error",
				Code:    http.StatusNotFound,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
		response.WriteJSON(w, http.StatusInternalServerError, response.Response{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Category updated successfully",
		Data:    category,
	})
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.Response{
			Status:  "error",
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	err = services.DeleteCategory(id)
	if err != nil {
		if _, ok := err.(*errors.ErrNotFound); ok {
			response.WriteJSON(w, http.StatusNotFound, response.Response{
				Status:  "error",
				Code:    http.StatusNotFound,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
		response.WriteJSON(w, http.StatusInternalServerError, response.Response{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Category deleted successfully",
		Data:    nil,
	})
}
