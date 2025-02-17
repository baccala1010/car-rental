package domain

type Feedback struct {
	ID          int64  `db:"id"`
	UserID      int64  `db:"user_id"`
	CarID       int64  `db:"car_id"`
	Rating      int    `db:"rating"`
	Description string `db:"description"`
}
