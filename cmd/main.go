package main

import (
	"log"

	"githhub.com/VSBrilyakov/test-app/internal"
	"githhub.com/VSBrilyakov/test-app/internal/handler"
	"githhub.com/VSBrilyakov/test-app/internal/repository"
	"githhub.com/VSBrilyakov/test-app/internal/service"
)

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(internal.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("server run failed: %s", err.Error())
	}
}
