package http

import (
	"net/http"
)

type ServerConfig struct {
	Source     string
	SourcePort string
}

func Init(config ServerConfig) {
	routes := make([]Route, 0)
	routes = append(routes, NewProductController())

	router := NewRouter(routes)
	router.CreateRoutes()

	http.ListenAndServe(":"+config.SourcePort, router.Get())
}
