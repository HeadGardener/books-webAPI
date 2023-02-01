package repository

type Authentication interface {
}

type BookInterface interface {
}

type Repository struct {
	Authentication
	BookInterface
}

func NewRepository() *Repository {
	return &Repository{}
}
