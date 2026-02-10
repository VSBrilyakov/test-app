package service

import (
	test_app "githhub.com/VSBrilyakov/test-app"
	"githhub.com/VSBrilyakov/test-app/internal/repository"
)

type SubscribeActions interface {
	CreateSubscription(sub test_app.Subscription) (int, error)
}
type Service struct {
	SubscribeActions
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		SubscribeActions: NewSubscriptionService(repo),
	}
}
