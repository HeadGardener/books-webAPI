package models

import "errors"

type Book struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Author      string `json:"author" db:"author" binding:"required"`
	Description string `json:"description" db:"description"`
}

type BookInput struct {
	Title       *string `json:"title"`
	Author      *string `json:"author"`
	Description *string `json:"description"`
}

func (bi BookInput) IsValid() error {
	if bi.Title == nil && bi.Author == nil && bi.Description == nil {
		return errors.New("empty update struct")
	}
	return nil
}
