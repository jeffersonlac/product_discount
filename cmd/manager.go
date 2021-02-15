package main

import (
	"product_discount/pkg/domain/product"
	"product_discount/pkg/http"
)

//BuildControllers build all controllers
func BuildControllers() []http.Route {
	controllers := make([]http.Route, 0)

	productRepository := product.NewTestRepository()
	productService := product.NewService(productRepository)
	productController := http.NewProductController(productService)
	controllers = append(controllers, productController)

	return controllers
}
