package product

type Repository interface {
	Reader
	Writer
}

type Reader interface {
	Find(id string) (*Product, error)
	FindAll() ([]*Product, error)
}

type Writer interface {
	Update(product *Product) error
	Store(product *Product) (string, error)
}
