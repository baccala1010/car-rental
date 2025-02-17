package handler

import (
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
	"gitlab.com/advanced-programing/car-rental-system/internal/mapper"
	"gitlab.com/advanced-programing/car-rental-system/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type RentalHandler struct {
	rentalService  service.RentalService
	paymentService service.PaymentService
}

func NewRentalHandler(rentalService service.RentalService, paymentService service.PaymentService) *RentalHandler {
	return &RentalHandler{rentalService: rentalService, paymentService: paymentService}
}

func (h *RentalHandler) RentCar(c *gin.Context) {
	var req struct {
		CarID     int64  `json:"car_id" binding:"required"`
		StartDate string `json:"start_date" binding:"required"`
		EndDate   string `json:"end_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, _ := time.Parse(time.RFC3339, req.StartDate)
	endDate, _ := time.Parse(time.RFC3339, req.EndDate)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	rental, err := h.rentalService.RentCar(userID.(int64), req.CarID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment := domain.Payment{
		RentalID: rental.ID,
		Status:   domain.PaymentPending,
	}
	createdPayment, err := h.paymentService.CreatePayment(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"rental":  mapper.RentalToDTO(rental),
		"payment": mapper.PaymentToDTO(createdPayment),
	})
}

func (h *RentalHandler) ReturnCar(c *gin.Context) {
	rentalIDStr := c.Param("id")
	rentalID, err := strconv.ParseInt(rentalIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid rental id"})
		return
	}
	err = h.rentalService.ReturnCar(rentalID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car returned successfully"})
}

func (h *RentalHandler) ListAllRentals(c *gin.Context) {
	rentals, err := h.rentalService.ListAllRentals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rentals": mapper.RentalsToDTOs(rentals)})
}

func (h *RentalHandler) ListRentalsByUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	rentals, err := h.rentalService.ListRentalsByUser(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rentals": mapper.RentalsToDTOs(rentals)})
}
