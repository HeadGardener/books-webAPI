package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createBook(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
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
