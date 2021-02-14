package http

import "github.com/go-chi/chi"

type Route interface {
	Route(r chi.Router)
	Path() string
}
