package http

import (
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router interface {
	CreateRoutes()
	Get() chi.Router
}

type Route interface {
	Route(r chi.Router)
	Path() string
}

type chiRouter struct {
	router chi.Router
	routes []Route
}

func NewRouter(routes []Route) Router {
	return chiRouter{
		router: chi.NewRouter(),
		routes: routes,
	}
}

func (c chiRouter) Get() chi.Router {
	return c.router
}

func (c chiRouter) CreateRoutes() {

	fmt.Printf("Creating routes\n")

	c.router.Route("/v1", func(r chi.Router) {
		r.Group(func(mainGroup chi.Router) {
			mainGroup.Use(middleware.Logger)
			mainGroup.Use(middleware.Recoverer)

			fmt.Printf("len c.routes: %d\n", len(c.routes))

			for _, handler := range c.routes {
				fmt.Printf("Registering: %s\n", handler.Path())
				mainGroup.Route(handler.Path(), handler.Route)
			}
		})
	})

	// c.Router.Get("/swagger/*", c.Swagger())
}
