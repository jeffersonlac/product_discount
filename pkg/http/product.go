package http

import (
	"net/http"
	"product_discount/pkg/domain/product"
	"product_discount/pkg/http/helpers"

	"github.com/go-chi/chi"
)

type productController struct {
	service *product.Service
}

//NewProductController create new product controller
func NewProductController(ps *product.Service) productController {
	return productController{
		service: ps,
	}
}

func (pc productController) Route(r chi.Router) {
	r.Get("/", pc.products)
	r.Get("/{id}", pc.product)
}

func (pc productController) Path() string {
	return "/products"
}

func (pc productController) products(w http.ResponseWriter, r *http.Request) {
	res, err := pc.service.List()
	if err != nil {
		helpers.JSONResponse(w, "Error to list products", http.StatusInternalServerError)
		return
	}
	helpers.JSONResponse(w, res, http.StatusOK)
}

func (pc productController) product(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, err := pc.service.Find(id)
	if err != nil {
		helpers.JSONResponse(w, "Error to find product: "+id, http.StatusInternalServerError)
		return
	}
	helpers.JSONResponse(w, res, http.StatusOK)
}
