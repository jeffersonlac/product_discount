package product

import "github.com/google/uuid"

type repo struct {
}

//NewTestRepository create new repository
func NewTestRepository() Repository {
	return &repo{}
}

func (r *repo) Find(id string) (*Product, error) {
	return &Product{
		ID:           id,
		PriceInCents: 100,
		Title:        "Produto 1",
		Description:  "Este é o produto 1",
	}, nil
}

func (r *repo) FindAll() ([]*Product, error) {
	products := make([]*Product, 0)
	p1, _ := r.Find(uuid.NewString())
	p2, _ := r.Find(uuid.NewString())
	products = append(products, p1, p2)

	return products, nil
}

func (r *repo) Update(product *Product) error {
	return nil
}

func (r *repo) Store(product *Product) (string, error) {
	return "", nil
}
