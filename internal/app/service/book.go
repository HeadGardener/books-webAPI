package service

import (
	"github.com/HeadGardener/books-webAPI/internal/app/models"
	"github.com/HeadGardener/books-webAPI/internal/app/repository"
)

type BookService struct {
	repos *repository.Repository
}

func NewBookService(repos *repository.Repository) *BookService {
	return &BookService{repos: repos}
}

func (s *BookService) CreateBook(userID int, book models.Book) (int, error) {
	return s.repos.BookInterface.CreateBook(userID, book)
}
