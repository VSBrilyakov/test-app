package service

import (
	test_app "githhub.com/VSBrilyakov/test-app"
	"githhub.com/VSBrilyakov/test-app/internal/repository"
)

type SubscriptionService struct {
	repo repository.SubscribeActions
}

func NewSubscriptionService(repo *repository.Repository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) CreateSubscription(sub test_app.Subscription) (int, error) {
	return s.repo.CreateSubscription(sub)
}

func (s *SubscriptionService) GetSubscription(subId int) (test_app.Subscription, error) {
	return s.repo.GetSubscription(subId)
}
