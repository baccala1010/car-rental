package handler

import (
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
	"gitlab.com/advanced-programing/car-rental-system/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	repo           interface{} // placeholder for user repository
	carService     service.CarService
	authService    service.AuthService
	rentalService  service.RentalService
	paymentService service.PaymentService
	userService    service.UserService
}

func NewAdminHandler(repo interface{}, carService service.CarService, authService service.AuthService, rentalService service.RentalService, paymentService service.PaymentService, userService service.UserService) *AdminHandler {
	return &AdminHandler{
		carService:     carService,
		authService:    authService,
		rentalService:  rentalService,
		paymentService: paymentService,
		userService:    userService,
	}
}

func (h *AdminHandler) CreateCar(c *gin.Context) {
	var car domain.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCar, err := h.carService.CreateCar(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdCar)
}

func (h *AdminHandler) UpdateCar(c *gin.Context) {
	var car domain.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car id"})
		return
	}
	car.ID = id
	err = h.carService.UpdateCar(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car updated"})
}

func (h *AdminHandler) DeleteCar(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car id"})
		return
	}
	err = h.carService.DeleteCar(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted"})
}

func (h *AdminHandler) ListAllUsers(c *gin.Context) {
	users, err := h.userService.ListAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
