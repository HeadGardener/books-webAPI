package handlers

import (
	"github.com/HeadGardener/books-webAPI/internal/app/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("api", h.identifyUser)
	{
		books := api.Group("/books")
		{
			books.POST("/", h.createBook)
			books.GET("/", h.getAllBooks)
			books.GET("/:id", h.getBookByID)
			books.PUT("/:id", h.updateBook)
			books.DELETE("/:id", h.deleteBook)
		}
	}

	return router
}
