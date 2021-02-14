package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router interface {
	CreateRoutes()
	Get() chi.Router
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
	c.router.Route("/v1", func(r chi.Router) {
		r.Group(func(mainGroup chi.Router) {
			mainGroup.Use(middleware.Logger)
			mainGroup.Use(middleware.Recoverer)

			for _, route := range c.routes {
				mainGroup.Route(route.Path(), route.Route)
			}
		})
	})

	// c.Router.Get("/swagger/*", c.Swagger())
}
