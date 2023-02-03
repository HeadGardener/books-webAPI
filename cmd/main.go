package main

import (
	"context"
	"github.com/HeadGardener/books-webAPI/internal/app/handlers"
	"github.com/HeadGardener/books-webAPI/internal/app/repository"
	"github.com/HeadGardener/books-webAPI/internal/app/service"
	"github.com/HeadGardener/books-webAPI/internal/pkg/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
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

	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("issue while running server: %s", err.Error())
		}
	}()

	logrus.Println("starting server...")

	osSignCh := make(chan os.Signal, 1)
	signal.Notify(osSignCh, syscall.SIGTERM, syscall.SIGINT)
	<-osSignCh

	logrus.Println("stopping server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error while server shutdown: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error while closing db connection: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
