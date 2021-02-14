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
	}

	return products, nil
}
