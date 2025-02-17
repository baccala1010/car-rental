package domain

import "time"

// UserRole enumeration
type UserRole string

const (
	RoleAdmin    UserRole = "ADMIN"
	RoleCustomer UserRole = "CUSTOMER"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Phone     string    `db:"phone"`
	Role      UserRole  `db:"role"`
	Password  string    `db:"password"` // stored as a hash
	CreatedAt time.Time `db:"created_at"`
}
