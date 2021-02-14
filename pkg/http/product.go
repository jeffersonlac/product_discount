package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

type productController struct {
	// service product.Service
}

func NewProductController() productController {
	//ps product.Service
	return productController{
		//service: ps,
	}
}

func (pc productController) Route(r chi.Router) {
	r.Get("/", pc.products)
}

func (pc productController) Path() string {
	return "/products"
}

func (pc productController) products(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
