package handler

import (
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
	"gitlab.com/advanced-programing/car-rental-system/internal/dto"
	"gitlab.com/advanced-programing/car-rental-system/internal/mapper"
	"gitlab.com/advanced-programing/car-rental-system/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FeedbackHandler struct {
	feedbackService service.FeedbackService
}

func NewFeedbackHandler(feedbackService service.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{feedbackService: feedbackService}
}

func (h *FeedbackHandler) CreateFeedback(c *gin.Context) {
	var req struct {
		CarID       int64  `json:"car_id" binding:"required"`
		Rating      int    `json:"rating" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	feedback := domain.Feedback{
		UserID:      userID.(int64),
		CarID:       req.CarID,
		Rating:      req.Rating,
		Description: req.Description,
	}
	createdFeedback, err := h.feedbackService.CreateFeedback(feedback)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapper.FeedbackToDTO(createdFeedback))
}

func (h *FeedbackHandler) ListFeedbackByCar(c *gin.Context) {
	carIDStr := c.Param("car_id")
	carID, err := strconv.ParseInt(carIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car id"})
		return
	}
	feedbacks, err := h.feedbackService.ListFeedbackByCar(carID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var feedbackDTOs []dto.FeedbackDTO
	for _, feedback := range feedbacks {
		feedbackDTOs = append(feedbackDTOs, mapper.FeedbackToDTO(feedback))
	}
	c.JSON(http.StatusOK, feedbackDTOs)
}

func (h *FeedbackHandler) ListFeedbackByUser(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	feedbacks, err := h.feedbackService.ListFeedbackByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var feedbackDTOs []dto.FeedbackDTO
	for _, feedback := range feedbacks {
		feedbackDTOs = append(feedbackDTOs, mapper.FeedbackToDTO(feedback))
	}
	c.JSON(http.StatusOK, feedbackDTOs)
}

func (h *FeedbackHandler) DeleteFeedbackAdmin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid feedback id"})
		return
	}
	err = h.feedbackService.DeleteFeedback(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Feedback deleted"})
}

func (h *FeedbackHandler) DeleteFeedbackClient(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid feedback id"})
		return
	}

	userID, _ := c.Get("userID")
	userRole, _ := c.Get("userRole")

	feedback, err := h.feedbackService.GetFeedback(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if feedback.UserID != userID.(int64) && userRole != domain.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized to delete this feedback"})
		return
	}

	if err := h.feedbackService.DeleteFeedback(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feedback deleted"})
}

func (h *FeedbackHandler) UpdateFeedback(c *gin.Context) {
	var feedback domain.Feedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid feedback id"})
		return
	}
	feedback.ID = id
	err = h.feedbackService.UpdateFeedback(feedback)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Feedback updated"})
}

func (h *FeedbackHandler) ListAllFeedback(c *gin.Context) {
	feedbacks, err := h.feedbackService.ListAllFeedback()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var feedbackDTOs []dto.FeedbackDTO
	for _, feedback := range feedbacks {
		feedbackDTOs = append(feedbackDTOs, mapper.FeedbackToDTO(feedback))
	}
	c.JSON(http.StatusOK, feedbackDTOs)
}
