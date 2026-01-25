package main

import (
	"encoding/json"
	"fmt"
	"kasir-api-bootcamp/services"
	"net/http"
)

func main() {
	fmt.Println("Server started in port 8080")

	// Products endpoint
	http.HandleFunc("GET /api/products", services.GetProducts)
	http.HandleFunc("GET /api/products/{id}", services.GetProduct)
	http.HandleFunc("POST /api/products", services.CreateProduct)
	http.HandleFunc("PUT /api/products/{id}", services.UpdateProduct)
	http.HandleFunc("DELETE /api/products/{id}", services.DeleteProduct)

	// Categories endpoint
	http.HandleFunc("GET /api/categories", services.GetCategories)
	http.HandleFunc("GET /api/categories/{id}", services.GetCategory)
	http.HandleFunc("POST /api/categories", services.CreateCategory)
	http.HandleFunc("PUT /api/categories/{id}", services.UpdateCategory)
	http.HandleFunc("DELETE /api/categories/{id}", services.DeleteCategory)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Server succesfuly running!",
			"data": map[string]interface{}{
				"app":     "Kasir API",
				"version": 1,
			},
		})
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
