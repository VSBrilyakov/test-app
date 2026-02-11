package test_app

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Subscription struct {
	ID          int          `json:"id,omitempty" db:"id"`
	ServiceName string       `json:"service_name" binding:"required" db:"service_name"`
	Price       int          `json:"price" binding:"required" db:"price"`
	UserID      string       `json:"user_id" binding:"required" db:"user_id"`
	StartDate   time.Time    `json:"start_date" binding:"required" time_format:"01-2006" db:"start_date"`
	EndDate     sql.NullTime `json:"end_date,omitempty" time_format:"01-2006" db:"end_date"`
}

type UpdSubscription struct {
	ServiceName *string       `json:"service_name" db:"service_name"`
	Price       *int          `json:"price" db:"price"`
	UserID      *string       `json:"user_id" db:"user_id"`
	StartDate   *sql.NullTime `json:"start_date" time_format:"01-2006" db:"start_date"`
	EndDate     *sql.NullTime `json:"end_date,omitempty" time_format:"01-2006" db:"end_date"`
}

func (u *UpdSubscription) UnmarshalJSON(data []byte) error {
	var temp SubscriptionJSON
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if temp.ServiceName == "" {
		u.ServiceName = nil
	} else {
		u.ServiceName = &temp.ServiceName
	}

	if temp.Price == 0 {
		u.Price = nil
	} else {
		u.Price = &temp.Price
	}

	if temp.UserID == "" {
		u.UserID = nil
	} else {
		u.UserID = &temp.UserID
	}

	if temp.StartDate != "" {
		parsedTime, err := time.Parse("01-2006", temp.StartDate)
		if err != nil {
			return err
		}

		nt := &sql.NullTime{
			Time:  parsedTime,
			Valid: true,
		}
		u.StartDate = nt
	} else {
		u.StartDate = nil
	}

	if temp.EndDate != "" {
		parsedTime, err := time.Parse("01-2006", temp.EndDate)
		if err != nil {
			return err
		}

		nt := &sql.NullTime{
			Time:  parsedTime,
			Valid: true,
		}
		u.EndDate = nt
	} else {
		u.EndDate = nil
	}

	return nil
}

type SubscriptionJSON struct {
	Id          int    `json:"id,omitempty" db:"id"`
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
		Id:          s.ID,
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

	s.ID = temp.Id
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

	subJSON.Id = s.ID
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
