package service

import (
	"errors"
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

func (s *BookService) GetAllBooks(userID int) ([]models.Book, error) {
	return s.repos.BookInterface.GetAllBooks(userID)
}

func (s *BookService) GetBookByID(userID, bookID int) (models.Book, error) {
	return s.repos.BookInterface.GetBookByID(userID, bookID)
}

func (s *BookService) UpdateBook(userID, bookID int, bookInput models.BookInput) error {
	_, err := s.GetBookByID(userID, bookID)
	if err != nil {
		return errors.New("book doesn't exist or you don't have enough rules to delete it")
	}

	return s.repos.BookInterface.UpdateBook(bookID, bookInput)
}

func (s *BookService) DeleteBook(userID, bookID int) error {
	_, err := s.GetBookByID(userID, bookID)
	if err != nil {
		return errors.New("book doesn't exist or you don't have enough rules to delete it")
	}

	return s.repos.BookInterface.DeleteBook(userID, bookID)
}
