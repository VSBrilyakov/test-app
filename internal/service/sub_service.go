package service

import (
	"time"

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

func (s *SubscriptionService) GetSubscription(subId int) (*test_app.Subscription, error) {
	return s.repo.GetSubscription(subId)
}

func (s *SubscriptionService) UpdateSubscription(subId int, input test_app.UpdSubscription) error {
	return s.repo.UpdateSubscription(subId, input)
}

func (s *SubscriptionService) DeleteSubscription(subId int) error {
	return s.repo.DeleteSubscription(subId)
}

func (s *SubscriptionService) GetAllSubscriptions() (*[]test_app.Subscription, error) {
	return s.repo.GetAllSubscriptions()
}

func (s *SubscriptionService) GetSubsSumByUserID(userId, serviceName string, dateFrom, dateTo time.Time) (int, error) {
	return s.repo.GetSubsSumByUserID(userId, serviceName, dateFrom, dateTo)
}
