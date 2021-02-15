package product

import "fmt"

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new product service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//List list all products
func (s Service) List() ([]*Product, error) {
	products, error := s.repo.FindAll()
	if error != nil {
		fmt.Printf("Error listing products: %v\n", error)
		return nil, error
	}

	return products, nil
}

//Find find the product by id
func (s Service) Find(id string) (*Product, error) {
	product, error := s.repo.Find(id)
	if error != nil {
		fmt.Printf("Error searching product: %v\n", error)
		return nil, error
	}

	return product, nil
}
