package repository

import (
	"fmt"
	"github.com/HeadGardener/books-webAPI/internal/app/models"
	"github.com/jmoiron/sqlx"
	"strings"
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
	row := tx.QueryRow(createBookQuery, book.Title, book.Author, book.Description)
	if err := row.Scan(&bookID); err != nil {
		tx.Rollback()
		return 0, err
	}

	usersBooksQuery := fmt.Sprintf("INSERT INTO %s (user_id, book_id) VALUES ($1, $2)", UsersBooksTable)
	_, err = tx.Exec(usersBooksQuery, userID, bookID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return bookID, tx.Commit()
}

func (r *BookPostgres) GetAllBooks(userID int) ([]models.Book, error) {
	var books []models.Book
	query := fmt.Sprintf(`SELECT b.id, b.title, b.author, b.description FROM %s
								b INNER JOIN %s ub on b.id = ub.book_id WHERE ub.user_id=$1`,
		BooksTable, UsersBooksTable)
	err := r.db.Select(&books, query, userID)

	return books, err
}

func (r *BookPostgres) GetBookByID(userID, bookID int) (models.Book, error) {
	var book models.Book
	query := fmt.Sprintf(`SELECT b.id, b.title, b.author, b.description FROM %s
								b INNER JOIN %s ub on b.id = ub.book_id WHERE ub.user_id=$1 AND b.id=$2`,
		BooksTable, UsersBooksTable)
	err := r.db.Get(&book, query, userID, bookID)

	return book, err
}

func (r *BookPostgres) UpdateBook(bookID int, bookInput models.BookInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if bookInput.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *bookInput.Title)
		argID++
	}

	if bookInput.Author != nil {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argID))
		args = append(args, *bookInput.Author)
		argID++
	}

	if bookInput.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *bookInput.Description)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s b SET %s WHERE b.id = $%d",
		BooksTable, setQuery, argID)
	args = append(args, bookID)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *BookPostgres) DeleteBook(userID, bookID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteFromUsersBooksQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND book_id=$2", UsersBooksTable)
	_, err = tx.Exec(deleteFromUsersBooksQuery, userID, bookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteFromBooksQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1", BooksTable)
	_, err = tx.Exec(deleteFromBooksQuery, bookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
