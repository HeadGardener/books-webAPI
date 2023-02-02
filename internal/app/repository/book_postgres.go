package repository

import (
	"fmt"
	"github.com/HeadGardener/books-webAPI/internal/app/models"
	"github.com/jmoiron/sqlx"
)

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) CreateBook(userID int, book models.Book) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var bookID int
	createBookQuery := fmt.Sprintf("INSERT INTO %s (title, author, description) VALUES ($1, $2, $3) RETURNING id",
		BooksTable)
	row := r.db.QueryRow(createBookQuery, book.Title, book.Author, book.Description)
	if err := row.Scan(&bookID); err != nil {
		tx.Rollback()
		return 0, err
	}

	usersBooksQuery := fmt.Sprintf("INSERT INTO %s (user_id, book_id) VALUES ($1, $2)", UsersBooksTable)
	_, err = r.db.Exec(usersBooksQuery, userID, bookID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return bookID, tx.Commit()
}
