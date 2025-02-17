package mapper

import (
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
	"gitlab.com/advanced-programing/car-rental-system/internal/dto"
)

func UserToDTO(user domain.User) dto.UserDTO {
	return dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Role:  user.Role,
	}
}

func CarToDTO(car domain.Car) dto.CarDTO {
	return dto.CarDTO{
		ID:           car.ID,
		Brand:        car.Brand,
		Model:        car.Model,
		Transmission: car.Transmission,
		FuelType:     car.FuelType,
		PricePerDay:  car.PricePerDay,
		Available:    car.Available,
		Photo:        car.Photo,
	}
}

func RentalToDTO(rental domain.Rental) dto.RentalDTO {
	return dto.RentalDTO{
		ID:        rental.ID,
		UserID:    rental.UserID,
		CarID:     rental.CarID,
		PaymentID: rental.PaymentID,
		StartDate: rental.StartDate,
		EndDate:   rental.EndDate,
	}
}

func PaymentToDTO(payment domain.Payment) dto.PaymentDTO {
	return dto.PaymentDTO{
		ID:       payment.ID,
		RentalID: payment.RentalID,
		Status:   payment.Status,
	}
}

func FeedbackToDTO(feedback domain.Feedback) dto.FeedbackDTO {
	return dto.FeedbackDTO{
		ID:          feedback.ID,
		UserID:      feedback.UserID,
		CarID:       feedback.CarID,
		Rating:      feedback.Rating,
		Description: feedback.Description,
	}
}

func RegisterRequestToUser(request dto.RegisterRequest) domain.User {
	return domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
	}
}

func LoginRequestToUser(request dto.LoginRequest) domain.User {
	return domain.User{
		Email:    request.Email,
		Password: request.Password,
	}
}

func RentalsToDTOs(rentals []domain.Rental) []dto.RentalDTO {
	var rentalDTOs []dto.RentalDTO
	for _, rental := range rentals {
		rentalDTOs = append(rentalDTOs, RentalToDTO(rental))
	}
	return rentalDTOs
}
