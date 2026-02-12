package main

import (
	"githhub.com/VSBrilyakov/test-app/configs"
	"githhub.com/VSBrilyakov/test-app/internal"
	"githhub.com/VSBrilyakov/test-app/internal/handler"
	"githhub.com/VSBrilyakov/test-app/internal/repository"
	"githhub.com/VSBrilyakov/test-app/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var config *configs.Config

func init() {
	var err error

	err = godotenv.Load()
	if err != nil {
		logrus.Fatal("invalid .env file")
	}

	config, err = configs.NewConfig()
	if err != nil {
		logrus.Fatalf("config reading error: %s", err.Error())
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
	})
	var logLvl logrus.Level
	if logLvl, err = logrus.ParseLevel(config.LogLevel); err != nil {
		logrus.Fatalf("invalid log level: %s", err.Error())
	}
	logrus.SetLevel(logLvl)
	logrus.Info("config loaded")

	gin.SetMode(gin.ReleaseMode)
}

func main() {
	db, err := repository.NewPostgresDB(&config.Postgres)
	if err != nil {
		logrus.Fatalf("postgres connection error: %s", err.Error())
	}
	logrus.Info("postgres connection established")

	err = repository.DoMigrates(db)
	if err != nil {
		logrus.Fatalf("migrations applying error: %s", err.Error())
	}
	logrus.Info("migrations have been applied")

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(internal.Server)
	if err := srv.Run(config.Server, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("server run failed: %s", err.Error())
	}
}
