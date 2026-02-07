package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"kasir-api/internal/database"
	"kasir-api/internal/handlers"
	"kasir-api/internal/repositories"
	"kasir-api/internal/router"
	"kasir-api/internal/services"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		fmt.Println("Gagal terhubung ke database", err)
	}

	// CRUD Product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Route Product : /api/product dan /api/product/
	router.ProductRegisterRoutes(productHandler)

	// CRUD Category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Route Category : /api/categories dan /api/categories/
	router.CategoryRegisterRoutes(categoryHandler)

	// Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Route Transaction
	router.TransactionRegisterRoutes(transactionHandler)

	// Health Check: http://localhost:8080/api/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Api Running",
		})
	})

	// Home
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		type Endpoint struct {
			Method      string `json:"method"`
			Path        string `json:"path"`
			Description string `json:"description"`
		}

		resp := map[string]any{
			"status":  "OK",
			"message": "Kasir API - endpoints",
			"endpoints": []Endpoint{
				{Method: "GET", Path: "/health", Description: "Health check"},
				{Method: "GET", Path: "/api/products", Description: "Get all products"},
				{Method: "POST", Path: "/api/products", Description: "Create product"},
				{Method: "GET", Path: "/api/products/{id}", Description: "Get product by ID"},
				{Method: "PUT", Path: "/api/products/{id}", Description: "Update product"},
				{Method: "DELETE", Path: "/api/products/{id}", Description: "Delete product"},
				{Method: "GET", Path: "/api/categories", Description: "Get all categories"},
				{Method: "POST", Path: "/api/categories", Description: "Create category"},
				{Method: "GET", Path: "/api/categories/{id}", Description: "Get category by ID"},
				{Method: "PUT", Path: "/api/categories/{id}", Description: "Update category"},
				{Method: "DELETE", Path: "/api/categories/{id}", Description: "Delete category"},
				{Method: "POST", Path: "/api/checkout", Description: "Create transaction (checkout)"},
				{Method: "GET", Path: "/api/report/hari-ini", Description: "Summary report for today"},
				{Method: "GET", Path: "/api/report?startDate=YYYY-MM-DD&endDate=YYYY-MM-DD", Description: "Summary report by date range"},
			},
		}

		json.NewEncoder(w).Encode(resp)
	})

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Gagal running server")
	}
}
