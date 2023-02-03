package handlers

import (
	"github.com/HeadGardener/books-webAPI/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	books, err := h.service.BookInterface.GetAllBooks(userID)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) getBookByID(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid book id")
		return
	}

	book, err := h.service.BookInterface.GetBookByID(userID, bookID)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) updateBook(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid book id")
		return
	}

	var bookInput models.BookInput
	if err := c.BindJSON(&bookInput); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid book data")
		return
	}

	if err := bookInput.IsValid(); err != nil {
		newErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.BookInterface.UpdateBook(userID, bookID, bookInput)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "updated",
	})
}

func (h *Handler) deleteBook(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid book id")
		return
	}

	err = h.service.BookInterface.DeleteBook(userID, bookID)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "deleted",
	})
}
