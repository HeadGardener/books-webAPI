package handlers

import (
	"github.com/HeadGardener/books-webAPI/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid data to create user")
		return
	}

	id, err := h.service.Authorization.CreateUser(user)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, "mistake creating user")
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.UserInput

	if err := c.BindJSON(&input); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid data for authorization user")
		return
	}

	token, err := h.service.Authorization.GenerateToken(input)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, "mistake in authorization")
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})
}
