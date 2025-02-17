package domain

import "time"

type Rental struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	CarID     int64     `db:"car_id"`
	PaymentID int64     `db:"payment_id"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
}
