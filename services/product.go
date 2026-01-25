package services

import (
	"encoding/json"
	"kasir-api-bootcamp/models"
	"kasir-api-bootcamp/repository"
	"net/http"
	"strconv"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: "Product list",
		Data:    repository.Products,
	})
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
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

	for _, p := range repository.Products {
		if p.ID == id {
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusOK,
				Message: "Products details",
				Data:    p,
			})
			return
		}
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNotFound,
		Message: "Product not found",
		Data:    nil,
	})
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	newProduct.ID = len(repository.Products) + 1
	repository.Products = append(repository.Products, newProduct)

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusCreated,
		Message: "Product added successfully",
		Data:    newProduct,
	})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
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

	var updatedProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	for i, p := range repository.Products {
		if p.ID == id {
			updatedProduct.ID = id
			repository.Products[i] = updatedProduct
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusOK,
				Message: "Product updated successfully",
				Data:    updatedProduct,
			})
			return
		}
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNotFound,
		Message: "Product not found",
		Data:    nil,
	})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
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

	for i, p := range repository.Products {
		if p.ID == id {
			deletedProduct := p
			repository.Products = append(repository.Products[:i], repository.Products[i+1:]...)
			json.NewEncoder(w).Encode(Response{
				Status:  http.StatusOK,
				Message: "Product deleted successfully",
				Data:    deletedProduct,
			})
			return
		}
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNotFound,
		Message: "Product not found",
		Data:    nil,
	})
}
