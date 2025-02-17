package domain

// PaymentStatus enum
type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "PENDING"
	PaymentCompleted PaymentStatus = "COMPLETED"
	PaymentFailed    PaymentStatus = "FAILED"
)

type Payment struct {
	ID       int64         `db:"id"`
	RentalID int64         `db:"rental_id"`
	Status   PaymentStatus `db:"status"`
}
