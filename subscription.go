package test_app

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Subscription struct {
	ID          int          `json:"-" db:"id"`
	ServiceName string       `json:"service_name" binding:"required" db:"service_name"`
	Price       int          `json:"price" binding:"required" db:"price"`
	UserID      string       `json:"user_id" binding:"required" db:"user_id"`
	StartDate   time.Time    `json:"start_date" binding:"required" time_format:"01-2006" db:"start_date"`
	EndDate     sql.NullTime `json:"end_date,omitempty" time_format:"01-2006" db:"end_date"`
}

type SubscriptionJSON struct {
	ServiceName string `json:"service_name" binding:"required" db:"service_name"`
	Price       int    `json:"price" binding:"required" db:"price"`
	UserID      string `json:"user_id" binding:"required" db:"user_id"`
	StartDate   string `json:"start_date" binding:"required" time_format:"01-2006" db:"start_date"`
	EndDate     string `json:"end_date,omitempty" time_format:"01-2006" db:"end_date"`
}

func (s *Subscription) MarshalJSON() ([]byte, error) {
	var endDate string
	if s.EndDate.Valid {
		endDate = s.EndDate.Time.Format("01-2006")
	} else {
		endDate = ""
	}

	return json.Marshal(SubscriptionJSON{
		ServiceName: s.ServiceName,
		Price:       s.Price,
		UserID:      s.UserID,
		StartDate:   s.StartDate.Format("01-2006"),
		EndDate:     endDate,
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

	parsedTime, err := time.Parse("01-2006", temp.StartDate)
	if err != nil {
		return err
	}
	s.StartDate = parsedTime
	s.EndDate.Valid = temp.EndDate != ""

	if s.EndDate.Valid {
		parsedTime, err = time.Parse("01-2006", temp.EndDate)
		if err != nil {
			return err
		}
		s.EndDate.Time = parsedTime
	}

	return nil
}

func (s *Subscription) GetJSON() SubscriptionJSON {
	var subJSON SubscriptionJSON

	subJSON.ServiceName = s.ServiceName
	subJSON.Price = s.Price
	subJSON.UserID = s.UserID
	subJSON.StartDate = s.StartDate.Format("01-2006")
	if s.EndDate.Valid {
		subJSON.EndDate = s.EndDate.Time.Format("01-2006")
	} else {
		subJSON.EndDate = ""
	}

	return subJSON
}
