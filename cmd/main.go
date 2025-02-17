package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"gitlab.com/advanced-programing/car-rental-system/internal/config"
	"gitlab.com/advanced-programing/car-rental-system/internal/handler"
	"gitlab.com/advanced-programing/car-rental-system/internal/repository"
	"gitlab.com/advanced-programing/car-rental-system/internal/service"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load configuration from environment variables.
	cfg := config.LoadConfig()

	// Build PostgreSQL connection string.
	dbConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Connect to PostgreSQL.
	db, err := sqlx.Connect("postgres", dbConnString)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Initialize the repository.
	repo := repository.NewPostgresRepository(db)

	// Initialize the real email service using SMTP configuration.
	emailService := service.NewEmailService(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPassword, cfg.SMTPFrom)
	authService := service.NewAuthService(repo, emailService, cfg.JWTSecret)
	carService := service.NewCarService(repo)
	rentalService := service.NewRentalService(repo, emailService)
	paymentService := service.NewPaymentService(repo)
	feedbackService := service.NewFeedbackService(repo)
	userService := service.NewUserService(repo)

	// Initialize HTTP handlers.
	authHandler := handler.NewAuthHandler(authService, cfg.AdminRegistrationSecret)
	carHandler := handler.NewCarHandler(carService)
	rentalHandler := handler.NewRentalHandler(rentalService, paymentService)
	adminHandler := handler.NewAdminHandler(repo, carService, authService, rentalService, paymentService, userService)
	feedbackHandler := handler.NewFeedbackHandler(feedbackService)
	userHandler := handler.NewUserHandler(userService)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	// Setup Gin router and public endpoints.
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Your React dev server
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Update static file serving to absolute path
	frontendDir := "./frontend/car-rental-frontend/build"
	r.Static("/static", filepath.Join(frontendDir, "static"))
	r.StaticFile("/manifest.json", filepath.Join(frontendDir, "manifest.json"))
	r.StaticFile("/logo192.png", filepath.Join(frontendDir, "logo192.png"))
	r.StaticFile("/logo512.png", filepath.Join(frontendDir, "logo512.png"))
	r.StaticFile("/robots.txt", filepath.Join(frontendDir, "robots.txt"))

	r.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(frontendDir, "index.html"))
	})
	// API endpoints
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "API is working"})
		})
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
		api.GET("/cars", carHandler.ListCars)
		api.GET("/cars/:id", carHandler.GetCar)
		api.GET("/cars/search", carHandler.ListCarsByCriteria)
		api.GET("/feedback/car/:car_id", feedbackHandler.ListFeedbackByCar)
		api.GET("/feedback/user/:user_id", feedbackHandler.ListFeedbackByUser)
		api.GET("/feedback", feedbackHandler.ListAllFeedback)

		// Protected endpoints (JWT required).
		authorized := api.Group("/")
		authorized.Use(handler.JWTMiddleware(cfg.JWTSecret))
		{
			authorized.POST("/rentals", rentalHandler.RentCar)
			authorized.POST("/rentals/:id/return", rentalHandler.ReturnCar)
			authorized.POST("/feedback", feedbackHandler.CreateFeedback)
			authorized.PUT("/user", userHandler.UpdateUser)
			authorized.DELETE("/feedback/:id", feedbackHandler.DeleteFeedbackClient)
			authorized.GET("/payments/:id", paymentHandler.GetPayment)
			authorized.POST("/logout", authHandler.Logout)
			authorized.GET("/rentals", rentalHandler.ListRentalsByUser)
		}

		// Admin endpoints (JWT + admin role).
		adminRoutes := api.Group("/admin")
		adminRoutes.Use(handler.JWTMiddleware(cfg.JWTSecret), handler.AdminMiddleware())
		{
			adminRoutes.POST("/cars", adminHandler.CreateCar)
			adminRoutes.PUT("/cars/:id", adminHandler.UpdateCar)
			adminRoutes.DELETE("/cars/:id", adminHandler.DeleteCar)
			adminRoutes.DELETE("/feedback/:id", feedbackHandler.DeleteFeedbackAdmin)
			adminRoutes.PUT("/payments/:id", paymentHandler.UpdatePayment)
			adminRoutes.DELETE("/payments/:id", paymentHandler.DeletePayment)
			adminRoutes.GET("/payments", paymentHandler.ListPayments)
			adminRoutes.GET("/rentals", rentalHandler.ListAllRentals)
		}
	}

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server running on port 8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
