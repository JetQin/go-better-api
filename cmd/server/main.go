package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"go-better-api/internal/database"
	"go-better-api/internal/handlers"

	appauth "go-better-api/internal/auth"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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

	api := r.Group("/api")
	api.GET("/users", handlers.GetUsers)
	api.POST("/users", handlers.CreateUser)
	api.GET("/users/:id", handlers.GetUser)
	api.PUT("/users/:id", handlers.UpdateUser)
	api.DELETE("/users/:id", handlers.DeleteUser)
	api.GET("/posts", handlers.GetPosts)
	api.POST("/posts", handlers.CreatePost)
	api.GET("/posts/:id", handlers.GetPost)
	api.PUT("/posts/:id", handlers.UpdatePost)
	api.DELETE("/posts/:id", handlers.DeletePost)

	// Start server
	if err := r.Run(addr); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
