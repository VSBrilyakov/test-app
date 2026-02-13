package repository

import (
	"time"

	test_app "github.com/VSBrilyakov/test-app"
	"github.com/jmoiron/sqlx"
)

type SubscribeActions interface {
	CreateSubscription(sub test_app.Subscription) (int, error)
	GetSubscription(subId int) (*test_app.Subscription, error)
	UpdateSubscription(subId int, input test_app.UpdSubscription) error
	DeleteSubscription(subId int) error
	GetAllSubscriptions() (*[]test_app.Subscription, error)
	GetSubsSumByUserID(userId, serviceName string, dateFrom, dateTo time.Time) (int, error)
}
type Repository struct {
	SubscribeActions
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SubscribeActions: NewSubPostgres(db),
	}
}
