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
	CreateBook(userID int, book models.Book) (int, error)
	GetAllBooks(userID int) ([]models.Book, error)
	GetBookByID(userID, bookID int) (models.Book, error)
	UpdateBook(userID, bookID int, bookInput models.BookInput) error
	DeleteBook(userID, bookID int) error
}

type Service struct {
	Authorization
	BookInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		BookInterface: NewBookService(repos),
	}
}
