package service

import (
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
	"gitlab.com/advanced-programing/car-rental-system/internal/repository"
)

type PaymentService interface {
	GetPayment(id int64) (domain.Payment, error)
	UpdatePayment(payment domain.Payment) error
	DeletePayment(id int64) error
	CreatePayment(payment domain.Payment) (domain.Payment, error)
	ListAllPayments() ([]domain.Payment, error)
}

type paymentService struct {
	repo repository.Repository
}

func NewPaymentService(repo repository.Repository) PaymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) GetPayment(id int64) (domain.Payment, error) {
	return s.repo.GetPaymentByID(id)
}

func (s *paymentService) UpdatePayment(payment domain.Payment) error {
	return s.repo.UpdatePayment(payment)
}

func (s *paymentService) DeletePayment(id int64) error {
	return s.repo.DeletePayment(id)
}

func (s *paymentService) CreatePayment(payment domain.Payment) (domain.Payment, error) {
	paymentID, err := s.repo.CreatePayment(payment)
	if err != nil {
		return domain.Payment{}, err
	}
	payment.ID = paymentID
	return payment, nil
}

func (s *paymentService) ListAllPayments() ([]domain.Payment, error) {
	return s.repo.ListAllPayments()
}
