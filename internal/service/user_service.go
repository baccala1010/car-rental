package service

import (
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
	"gitlab.com/advanced-programing/car-rental-system/internal/repository"
)

type UserService interface {
	UpdateUser(user domain.User) error
	ListAllUsers() ([]domain.User, error)
	ListUserByEmail(email string) ([]domain.User, error)
	GetUserByID(id int) (*domain.User, error)
}

type userService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{repo: repo}
}

func (s *userService) UpdateUser(user domain.User) error {
	// Fetch the existing user to retain the role
	existingUser, err := s.repo.GetUserByID(user.ID)
	if err != nil {
		return err
	}
	user.Role = existingUser.Role
	user.Password = existingUser.Password
	return s.repo.UpdateUser(user)
}

func (s *userService) ListAllUsers() ([]domain.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) ListUserByEmail(email string) ([]domain.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return []domain.User{user}, nil
}

func (s *userService) GetUserByID(id int) (*domain.User, error) {
	return s.repo.GetUserByID(int64(id))
}
