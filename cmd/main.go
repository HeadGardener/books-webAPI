package main

import (
	"github.com/HeadGardener/books-webAPI/internal/app/handlers"
	"github.com/HeadGardener/books-webAPI/internal/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	handler := &handlers.Handler{}
	srv := &server.Server{}
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		logrus.Fatalf("issue while running server: %s", err.Error())
	}
}
