package service

import (
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
	"gitlab.com/advanced-programing/car-rental-system/internal/repository"
)

type CarService interface {
	ListCars() ([]domain.Car, error)
	CreateCar(car domain.Car) (domain.Car, error)
	UpdateCar(car domain.Car) error
	DeleteCar(id int64) error
	GetCar(id int64) (domain.Car, error)
	ListCarsByCriteria(criteria map[string]interface{}) ([]domain.Car, error)
}

type carService struct {
	repo repository.Repository
}

func NewCarService(repo repository.Repository) CarService {
	return &carService{repo: repo}
}

func (s *carService) ListCars() ([]domain.Car, error) {
	return s.repo.ListCars()
}

func (s *carService) CreateCar(car domain.Car) (domain.Car, error) {
	id, err := s.repo.CreateCar(car)
	if err != nil {
		return domain.Car{}, err
	}
	car.ID = id
	return car, nil
}

func (s *carService) UpdateCar(car domain.Car) error {
	return s.repo.UpdateCar(car)
}

func (s *carService) DeleteCar(id int64) error {
	return s.repo.DeleteCar(id)
}

func (s *carService) GetCar(id int64) (domain.Car, error) {
	return s.repo.GetCarByID(id)
}

func (s *carService) ListCarsByCriteria(criteria map[string]interface{}) ([]domain.Car, error) {
	return s.repo.ListCarsByCriteria(criteria)
}
