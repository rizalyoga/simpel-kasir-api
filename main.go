package main

import (
	"encoding/json"
	"fmt"
	"kasir-api-bootcamp/database"
	"kasir-api-bootcamp/handlers"
	"kasir-api-bootcamp/repositories"
	"kasir-api-bootcamp/services"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.SetEnvPrefix("")
	viper.AutomaticEnv()

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	fmt.Println("Server started in port", config.Port)

	http.HandleFunc("GET /api/products", productHandler.HandleProducts)
	http.HandleFunc("POST /api/products", productHandler.HandleProducts)
	http.HandleFunc("GET /api/products/", productHandler.HandleProductByID)
	http.HandleFunc("PUT /api/products/", productHandler.HandleProductByID)
	http.HandleFunc("DELETE /api/products/", productHandler.HandleProductByID)

	http.HandleFunc("GET /api/categories", handlers.GetCategories)
	http.HandleFunc("GET /api/categories/{id}", handlers.GetCategory)
	http.HandleFunc("POST /api/categories", handlers.CreateCategory)
	http.HandleFunc("PUT /api/categories/{id}", handlers.UpdateCategory)
	http.HandleFunc("DELETE /api/categories/{id}", handlers.DeleteCategory)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	if config.Port == "" {
		config.Port = "8080"
	}
	err = http.ListenAndServe(":"+config.Port, nil)

	if err != nil {
		fmt.Println("Failed to run server:", err)
	}
}
