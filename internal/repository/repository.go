package repository

import (
	test_app "githhub.com/VSBrilyakov/test-app"
	"github.com/jmoiron/sqlx"
)

type SubscribeActions interface {
	CreateSubscription(sub test_app.Subscription) (int, error)
}
type Repository struct {
	SubscribeActions
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SubscribeActions: NewSubPostgres(db),
	}
}
