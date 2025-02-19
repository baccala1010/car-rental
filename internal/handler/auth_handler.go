// internal/handler/auth_handler.go
package handler

import (
	"net/http"
	"strings"

	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
	"gitlab.com/advanced-programing/car-rental-system/internal/dto"
	"gitlab.com/advanced-programing/car-rental-system/internal/mapper"
	"gitlab.com/advanced-programing/car-rental-system/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
	adminSecret string
}

func NewAuthHandler(authService service.AuthService, adminSecret string) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		adminSecret: adminSecret,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set role based on whether the admin secret is provided.
	role := domain.RoleCustomer
	if req.AdminSecret != "" && req.AdminSecret == h.adminSecret {
		role = domain.RoleAdmin
	}

	user, err := h.authService.Register(req.Name, req.Email, req.Phone, req.Password, role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, mapper.UserToDTO(user))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// LogoutHandler handles user logout.
func (h *AuthHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header required"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Blacklist the token (implement BlacklistToken accordingly)
	if err := BlacklistToken(tokenString); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not log out"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
