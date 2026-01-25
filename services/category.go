package services

import (
	"encoding/json"
	"kasir-api-bootcamp/models"
	"kasir-api-bootcamp/repository"
	"net/http"
	"strconv"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: "Category list",
		Data:    repository.Categories,
	})
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	for _, c := range repository.Categories {
		if c.ID == id {
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusOK,
				Message: "Category details",
				Data:    c,
			})
			return
		}
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNotFound,
		Message: "Category not found",
		Data:    nil,
	})
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCategory models.Category
	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	newCategory.ID = len(repository.Categories) + 1
	repository.Categories = append(repository.Categories, newCategory)

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusCreated,
		Message: "Category added successfully",
		Data:    newCategory,
	})
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	var updatedCategory models.Category
	if err := json.NewDecoder(r.Body).Decode(&updatedCategory); err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	for i, p := range repository.Categories {
		if p.ID == id {
			updatedCategory.ID = id
			repository.Categories[i] = updatedCategory
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusOK,
				Message: "Category updated successfully",
				Data:    updatedCategory,
			})
			return
		}
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNotFound,
		Message: "Category not found",
		Data:    nil,
	})
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	for i, c := range repository.Categories {
		if c.ID == id {
			deletedCategory := c
			repository.Categories = append(repository.Categories[:i], repository.Categories[i+1:]...)
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusOK,
				Message: "Category deleted successfully",
				Data:    deletedCategory,
			})
			return
		}
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNotFound,
		Message: "Category not found",
		Data:    nil,
	})
}
