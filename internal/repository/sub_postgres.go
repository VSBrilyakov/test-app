package repository

import (
	"fmt"

	test_app "githhub.com/VSBrilyakov/test-app"
	"github.com/jmoiron/sqlx"
)

type SubPostgres struct {
	db *sqlx.DB
}

func NewSubPostgres(db *sqlx.DB) *SubPostgres {
	return &SubPostgres{db: db}
}

func (s *SubPostgres) CreateSubscription(sub test_app.Subscription) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (service_name, price, user_id, start_date, end_date) VALUES ($1, $2, $3, $4, $5) RETURNING id", subscriptionTable)

	row := s.db.QueryRow(query, sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, sub.EndDate)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SubPostgres) GetSubscription(subId int) (test_app.Subscription, error) {
	var subscription test_app.Subscription

	query := fmt.Sprintf("SELECT id, service_name, price, user_id, start_date, end_date FROM %s WHERE id = $1", subscriptionTable)
	err := s.db.Get(&subscription, query, subId)

	return subscription, err
}
