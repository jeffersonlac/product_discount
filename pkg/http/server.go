package http

import (
	"net/http"
)

//ServerConfig is the server data config
type ServerConfig struct {
	Source     string
	SourcePort string
}

//Init start the http server
func Init(config ServerConfig, controllers []Route) {

	router := NewRouter(controllers)
	router.CreateRoutes()

	http.ListenAndServe(":"+config.SourcePort, router.Get())
}
