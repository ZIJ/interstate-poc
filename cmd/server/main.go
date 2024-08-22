package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zij/interstate/handlers"
	"github.com/zij/interstate/pkg/s3client"
	"github.com/zij/interstate/service"
)

func main() {
	// Initialize S3 client
	s3Client, err := s3client.NewS3Client()
	if err != nil {
		log.Fatalf("Failed to initialize S3 client: %v", err)
	}

	// Initialize backend service
	backendService := service.NewBackendService(s3Client)

	// Initialize backend handler
	backendHandler := handlers.NewBackendHandler(backendService)

	// Set up Gin router
	router := gin.Default()

	// Define API routes
	api := router.Group("/api")
	{
		backends := api.Group("/backends")
		{
			backends.GET("", backendHandler.ListBackends)
			backends.POST("", backendHandler.CreateBackend)
			backends.GET("/:backendId", backendHandler.GetBackend)
			backends.PUT("/:backendId", backendHandler.UpdateBackend)
			backends.DELETE("/:backendId", backendHandler.DeleteBackend)

			backends.GET("/:backendId/state", backendHandler.GetBackendState)
			backends.POST("/:backendId/state", backendHandler.UpdateBackendState)
			backends.DELETE("/:backendId/state", backendHandler.ResetBackendState)
		}
	}

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Starting server on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}