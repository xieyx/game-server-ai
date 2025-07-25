package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xieyx/game-server-ai/internal/database"
	"github.com/xieyx/game-server-ai/internal/handlers"
	"github.com/xieyx/game-server-ai/internal/services"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	// Connect to database
	database.ConnectDB()
	database.MigrateDB()

	// Initialize services
	userService := services.NewUserService()

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)

	// Set up routes
	router := gin.Default()
	setupRoutes(router, userHandler)

	// Get server port from environment variable
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server starting on port %s", port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func setupRoutes(router *gin.Engine, userHandler handlers.UserHandlerInterface) {
	// User routes
	router.POST("/users", userHandler.CreateUser)
	router.POST("/login", userHandler.Login)
	router.GET("/users/:id", userHandler.GetUser)

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}
