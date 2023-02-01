package main

import (
	"github.com/HeadGardener/books-webAPI/internal/app/handlers"
	"github.com/HeadGardener/books-webAPI/internal/app/repository"
	"github.com/HeadGardener/books-webAPI/internal/app/service"
	"github.com/HeadGardener/books-webAPI/internal/pkg/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error occured while initialize config: %s", err.Error())
	}

	db, err := repository.NewDB(repository.Config{
		Host:    viper.GetString("db.host"),
		DBName:  viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failing while connecting to db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	srv := &server.Server{}
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("issue while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
