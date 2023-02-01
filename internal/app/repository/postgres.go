package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	UsersTable      = "users"
	BooksTable      = "books"
	UsersBooksTable = "users_books"
)

type Config struct {
	Host    string
	DBName  string
	SSLMode string
}

func NewDB(conf Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s dbname=%s sslmode=%s", conf.Host, conf.DBName, conf.SSLMode))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
