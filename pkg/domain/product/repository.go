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
	Update(user *Product) error
	Store(user *Product) (string, error)
}
