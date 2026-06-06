package main

import (
	"os"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-better-api/internal/database"

	appauth "go-better-api/internal/auth"

	"github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
)

var logger = logrus.New()

func initLogger() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	if os.Getenv("LOG_LEVEL") == "debug" {
		logger.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	// Load environment variables
	godotenv.Load()

	// Initialize logger
	initLogger()

	logger.Info("Starting go-better-api server")



	// JWT auth helper available via appauth.GenerateToken / appauth.ValidateToken
	_ = appauth.ErrInvalidToken // ensure auth package is imported

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	_ = db // Use db in your handlers

	// Get host from environment
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	// Get HTTP port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := host + ":" + port


	// Create Gin router
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})


	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// Root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to go-better-api!",
		})
	})

	// Start server
	if err := r.Run(addr); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
