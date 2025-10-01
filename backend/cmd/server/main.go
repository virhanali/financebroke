package main

import (
	"log"
	"os"
	"finance-app/backend/internal/database"
	"finance-app/backend/internal/handler"
	"finance-app/backend/internal/middleware"
	"finance-app/backend/internal/repository"
	"finance-app/backend/internal/services"
	"finance-app/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	database.Connect()

	// Initialize services
	telegramService := services.NewTelegramService(os.Getenv("TELEGRAM_BOT_TOKEN"))
	emailService := services.NewEmailService(
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("FROM_EMAIL"),
	)

	// Initialize repositories
	userRepo := repository.NewUserRepository(database.DB)
	billRepo := repository.NewBillRepository(database.DB)

	// Initialize usecases
	authUsecase := usecase.NewAuthUsecase(userRepo)
	billUsecase := usecase.NewBillUsecase(billRepo)
	notificationUsecase := usecase.NewNotificationUsecase(userRepo, telegramService, emailService)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authUsecase)
	billHandler := handler.NewBillHandler(billUsecase)
	dashboardHandler := handler.NewDashboardHandler(billUsecase)
	notificationHandler := handler.NewNotificationHandler(notificationUsecase)

	// Setup router
	r := gin.Default()

	// Logging middleware
	r.Use(middleware.LoggingMiddleware())

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Public routes
	public := r.Group("/api/v1")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
	}

	// Protected routes
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		// Auth
		protected.GET("/profile", authHandler.GetProfile)

		// Bills
		protected.GET("/bills", billHandler.GetBills)
		protected.POST("/bills", billHandler.CreateBill)
		protected.GET("/bills/:id", billHandler.GetBill)
		protected.PUT("/bills/:id", billHandler.UpdateBill)
		protected.DELETE("/bills/:id", billHandler.DeleteBill)
		protected.GET("/bills/upcoming", billHandler.GetUpcomingBills)

		// Dashboard
		protected.GET("/dashboard", dashboardHandler.GetDashboard)

		// Notifications
		protected.PUT("/notifications/settings", notificationHandler.UpdateNotificationSettings)
		protected.POST("/notifications/test-telegram", notificationHandler.TestTelegram)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}