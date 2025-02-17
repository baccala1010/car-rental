package handler

import (
	"gitlab.com/advanced-programing/car-rental-system/internal/mapper"
	"gitlab.com/advanced-programing/car-rental-system/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	carService service.CarService
}

func NewCarHandler(carService service.CarService) *CarHandler {
	return &CarHandler{carService: carService}
}

func (h *CarHandler) ListCars(c *gin.Context) {
	cars, err := h.carService.ListCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var carDTOs []interface{}
	for _, car := range cars {
		carDTOs = append(carDTOs, mapper.CarToDTO(car))
	}
	c.JSON(http.StatusOK, carDTOs)
}

func (h *CarHandler) ListCarsByCriteria(c *gin.Context) {
	criteria := make(map[string]interface{})
	for key, values := range c.Request.URL.Query() {
		criteria[key] = values
	}
	cars, err := h.carService.ListCarsByCriteria(criteria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var carDTOs []interface{}
	for _, car := range cars {
		carDTOs = append(carDTOs, mapper.CarToDTO(car))
	}
	c.JSON(http.StatusOK, carDTOs)
}

func (h *CarHandler) GetCar(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car id"})
		return
	}
	car, err := h.carService.GetCar(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mapper.CarToDTO(car))
}
