package main

import (
	"fmt"

	"product_discount/pkg/domain/product"
	"product_discount/pkg/http"
)

func main() {
	fmt.Println("Starting...")

	serverConfig := http.ServerConfig{
		Source:     "localhost",
		SourcePort: "8000",
	}

	repository := product.NewTestRepository()
	productService := product.NewService(repository)
	userController := http.NewProductController(productService)

	controllers := make([]http.Route, 0)
	controllers = append(controllers, userController)

	http.Init(serverConfig, controllers)
}
