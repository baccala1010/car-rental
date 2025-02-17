package service

import (
	"errors"
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
	"gitlab.com/advanced-programing/car-rental-system/internal/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(name, email, phone, password string, role domain.UserRole) (domain.User, error)
	Login(email, password string) (string, error)
}

type authService struct {
	repo         repository.Repository
	emailService EmailService
	jwtSecret    string
}

func NewAuthService(repo repository.Repository, emailService EmailService, jwtSecret string) AuthService {
	return &authService{repo: repo, emailService: emailService, jwtSecret: jwtSecret}
}

func (s *authService) Register(name, email, phone, password string, role domain.UserRole) (domain.User, error) {
	// Check if user exists
	_, err := s.repo.GetUserByEmail(email)
	if err == nil {
		return domain.User{}, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}
	newUser := domain.User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Role:     role,
		Password: string(hashedPassword),
	}
	id, err := s.repo.CreateUser(newUser)
	if err != nil {
		return domain.User{}, err
	}
	newUser.ID = id

	// Send welcome email
	s.emailService.SendEmail(email, "Welcome!", "Thank you for registering!")
	return newUser, nil
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}
	// Send login notification email
	s.emailService.SendEmail(user.Email, "Login Alert", "You have successfully logged in.")
	return tokenString, nil
}
