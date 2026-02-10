package test_app

import (
	"encoding/json"
	"time"
)

type Subscription struct {
	ServiceName string    `json:"service_name" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	UserID      string    `json:"user_id" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date"`
}

type SubscriptionJSON struct {
	ServiceName string `json:"service_name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date"`
}

func (s *Subscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(SubscriptionJSON{
		ServiceName: s.ServiceName,
		Price:       s.Price,
		UserID:      s.UserID,
		StartDate:   s.StartDate.Format("2006-01-02"),
		EndDate:     s.EndDate.Format("2006-01-02"),
	})
}

func (s *Subscription) UnmarshalJSON(data []byte) error {
	var temp SubscriptionJSON
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	s.ServiceName = temp.ServiceName
	s.Price = temp.Price
	s.UserID = temp.UserID

	parsedTime, err := time.Parse("2006-01-02", temp.StartDate)
	if err != nil {
		return err
	}
	s.StartDate = parsedTime

	parsedTime, err = time.Parse("2006-01-02", temp.EndDate)
	if err != nil {
		return err
	}
	s.EndDate = parsedTime

	return nil
}
