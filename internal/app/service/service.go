package service

import (
	"github.com/HeadGardener/books-webAPI/internal/app/models"
	"github.com/HeadGardener/books-webAPI/internal/app/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(inputUser models.UserInput) (string, error)
	ParseToken(token string) (int, error)
}

type BookInterface interface {
}

type Service struct {
	Authorization
	BookInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
	}
}
