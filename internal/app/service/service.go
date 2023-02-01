package service

import "github.com/HeadGardener/books-webAPI/internal/app/repository"

type Authentication interface {
}

type BookInterface interface {
}

type Service struct {
	Authentication
	BookInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
