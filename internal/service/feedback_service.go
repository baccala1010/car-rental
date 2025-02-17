package service

import (
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
	"gitlab.com/advanced-programing/car-rental-system/internal/repository"
)

type FeedbackService interface {
	CreateFeedback(feedback domain.Feedback) (domain.Feedback, error)
	UpdateFeedback(feedback domain.Feedback) error
	DeleteFeedback(id int64) error
	ListFeedbackByCar(carID int64) ([]domain.Feedback, error)
	ListFeedbackByUser(userID int64) ([]domain.Feedback, error)
	ListAllFeedback() ([]domain.Feedback, error)
	GetFeedback(id int64) (domain.Feedback, error)
}

type feedbackService struct {
	repo repository.Repository
}

func NewFeedbackService(repo repository.Repository) FeedbackService {
	return &feedbackService{repo: repo}
}

func (s *feedbackService) GetFeedback(id int64) (domain.Feedback, error) {
	return s.repo.GetFeedbackByID(id)
}

func (s *feedbackService) CreateFeedback(feedback domain.Feedback) (domain.Feedback, error) {
	feedbackID, err := s.repo.CreateFeedback(feedback)
	if err != nil {
		return domain.Feedback{}, err
	}
	feedback.ID = feedbackID
	return feedback, nil
}

func (s *feedbackService) UpdateFeedback(feedback domain.Feedback) error {
	return s.repo.UpdateFeedback(feedback)
}

func (s *feedbackService) DeleteFeedback(id int64) error {
	return s.repo.DeleteFeedback(id)
}

func (s *feedbackService) ListFeedbackByCar(carID int64) ([]domain.Feedback, error) {
	return s.repo.ListFeedbackByCar(carID)
}

func (s *feedbackService) ListFeedbackByUser(userID int64) ([]domain.Feedback, error) {
	return s.repo.ListFeedbackByUser(userID)
}

func (s *feedbackService) ListAllFeedback() ([]domain.Feedback, error) {
	return s.repo.ListAllFeedback()
}
