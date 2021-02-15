package main

import (
	"fmt"

	"product_discount/pkg/http"
)

func main() {
	fmt.Println("Starting...")

	serverConfig := http.ServerConfig{
		Source:     "localhost",
		SourcePort: "8000",
	}

	controllers := BuildControllers()

	http.Init(serverConfig, controllers)
}
