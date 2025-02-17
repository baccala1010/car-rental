package service

import (
	"errors"
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
	"gitlab.com/advanced-programing/car-rental-system/internal/repository"
	"time"
)

type RentalService interface {
	RentCar(userID, carID int64, startDate, endDate time.Time) (domain.Rental, error)
	ReturnCar(rentalID int64) error
	ListAllRentals() ([]domain.Rental, error)
	ListRentalsByUser(userID int64) ([]domain.Rental, error)
}

type rentalService struct {
	repo         repository.Repository
	emailService EmailService
}

func NewRentalService(repo repository.Repository, emailService EmailService) RentalService {
	return &rentalService{repo: repo, emailService: emailService}
}

func (s *rentalService) RentCar(userID, carID int64, startDate, endDate time.Time) (domain.Rental, error) {
	// Check if the car is available
	car, err := s.repo.GetCarByID(carID)
	if err != nil {
		return domain.Rental{}, err
	}
	if !car.Available {
		return domain.Rental{}, errors.New("car is not available")
	}

	// Create the rental record
	rental := domain.Rental{
		UserID:    userID,
		CarID:     carID,
		StartDate: startDate,
		EndDate:   endDate,
	}
	rentalID, err := s.repo.CreateRental(rental)
	if err != nil {
		return domain.Rental{}, err
	}
	rental.ID = rentalID

	// Create a payment record with the correct rental ID
	payment := domain.Payment{
		RentalID: rentalID,
		Status:   domain.PaymentPending,
	}
	paymentID, err := s.repo.CreatePayment(payment)
	if err != nil {
		return domain.Rental{}, err
	}
	rental.PaymentID = paymentID

	// Update the rental record with the payment ID
	err = s.repo.UpdateRental(rental)
	if err != nil {
		return domain.Rental{}, err
	}

	// Mark the car as unavailable
	car.Available = false
	s.repo.UpdateCar(car)

	// (For brevity, we omit fetching the user details here.)
	s.emailService.SendEmail("", "Car Rental Confirmation", "You have successfully rented the car.")
	return rental, nil
}

func (s *rentalService) ReturnCar(rentalID int64) error {
	rental, err := s.repo.GetRentalByID(rentalID)
	if err != nil {
		return err
	}
	// Mark the car as available
	car, err := s.repo.GetCarByID(rental.CarID)
	if err != nil {
		return err
	}
	car.Available = true
	s.repo.UpdateCar(car)

	// Finalize payment – if returned late, extra charges might apply (simplified logic)
	payment, err := s.repo.GetPaymentByID(rental.PaymentID)
	if err != nil {
		return err
	}
	if time.Now().After(rental.EndDate) {
		// Apply additional fees if needed (logic omitted for brevity)
		payment.Status = domain.PaymentCompleted
	} else {
		payment.Status = domain.PaymentCompleted
	}
	s.repo.UpdatePayment(payment)
	// Optionally send an email notification about return confirmation.
	return nil
}

func (s *rentalService) ListAllRentals() ([]domain.Rental, error) {
	return s.repo.GetAllRentals()
}

func (s *rentalService) ListRentalsByUser(userID int64) ([]domain.Rental, error) {
	return s.repo.GetRentalsByUser(userID)
}
