package handlers

import (
	"github.com/HeadGardener/books-webAPI/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createBook(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var bookInput models.Book
	if err := c.BindJSON(&bookInput); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid book data")
		return
	}

	bookID, err := h.service.BookInterface.CreateBook(userID, bookInput)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"bookID": bookID,
	})
}

func (h *Handler) getAllBooks(c *gin.Context) {

}

func (h *Handler) getBookByID(c *gin.Context) {

}

func (h *Handler) updateBook(c *gin.Context) {

}

func (h *Handler) deleteBook(c *gin.Context) {

}
