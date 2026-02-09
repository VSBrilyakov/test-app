package service

import "githhub.com/VSBrilyakov/test-app/internal/repository"

type SubscribeActions interface {
}

type Service struct {
	SubscribeActions
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
