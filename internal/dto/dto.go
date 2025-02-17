package dto

import (
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
	"time"
)

type UserDTO struct {
	ID    int64           `json:"id"`
	Name  string          `json:"name"`
	Email string          `json:"email"`
	Phone string          `json:"phone"`
	Role  domain.UserRole `json:"role"`
}

type RegisterRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Phone       string `json:"phone"`
	Password    string `json:"password" binding:"required"`
	AdminSecret string `json:"admin_secret"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CarDTO struct {
	ID           int64               `json:"id"`
	Brand        domain.CarBrand     `json:"brand"`
	Model        domain.CarModel     `json:"model"`
	Transmission domain.Transmission `json:"transmission"`
	FuelType     domain.FuelType     `json:"fuel_type"`
	PricePerDay  float64             `json:"price_per_day"`
	Available    bool                `json:"available"`
	Photo        string              `json:"photo"`
}

type RentalDTO struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CarID     int64     `json:"car_id"`
	PaymentID int64     `json:"payment_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type PaymentDTO struct {
	ID       int64                `json:"id"`
	RentalID int64                `json:"rental_id"`
	Status   domain.PaymentStatus `json:"status"`
}

type FeedbackDTO struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	CarID       int64  `json:"car_id"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}
