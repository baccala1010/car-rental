package repository

import (
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
)

type Repository interface {
	// User methods
	CreateUser(user domain.User) (int64, error)
	GetUserByEmail(email string) (domain.User, error)
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id int64) (*domain.User, error)
	UpdateUser(user domain.User) error

	// Car methods
	CreateCar(car domain.Car) (int64, error)
	UpdateCar(car domain.Car) error
	DeleteCar(id int64) error
	GetCarByID(id int64) (domain.Car, error)
	ListCars() ([]domain.Car, error)
	ListCarsByCriteria(criteria map[string]interface{}) ([]domain.Car, error)

	// Rental methods
	CreateRental(rental domain.Rental) (int64, error)
	UpdateRental(rental domain.Rental) error
	GetRentalByID(id int64) (domain.Rental, error)
	GetAllRentals() ([]domain.Rental, error)
	GetRentalsByUser(userID int64) ([]domain.Rental, error)

	// Payment methods
	CreatePayment(payment domain.Payment) (int64, error)
	UpdatePayment(payment domain.Payment) error
	GetPaymentByID(id int64) (domain.Payment, error)
	DeletePayment(id int64) error
	ListAllPayments() ([]domain.Payment, error)

	// Feedback methods
	CreateFeedback(feedback domain.Feedback) (int64, error)
	UpdateFeedback(feedback domain.Feedback) error
	DeleteFeedback(id int64) error
	GetFeedbackByID(id int64) (domain.Feedback, error)
	ListFeedbackByCar(carID int64) ([]domain.Feedback, error)
	ListFeedbackByUser(userID int64) ([]domain.Feedback, error)
	ListAllFeedback() ([]domain.Feedback, error)
}
