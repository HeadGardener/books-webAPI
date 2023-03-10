package repository

import (
	"github.com/HeadGardener/books-webAPI/internal/app/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(userInput models.UserInput) (models.User, error)
}

type BookInterface interface {
	CreateBook(userID int, book models.Book) (int, error)
	GetAllBooks(userID int) ([]models.Book, error)
	GetBookByID(userID, bookID int) (models.Book, error)
	UpdateBook(bookID int, bookInput models.BookInput) error
	DeleteBook(userID, bookID int) error
}

type Repository struct {
	Authorization
	BookInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		BookInterface: NewBookPostgres(db),
	}
}
