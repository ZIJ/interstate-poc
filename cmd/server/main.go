package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zij/interstate/handlers"
)

func main() {
	r := gin.Default()

	// Create a new BackendHandler
	backendHandler := handlers.NewBackendHandler()

	// Group all routes under /api
	api := r.Group("/api")

	// Backends routes
	backends := api.Group("/backends")
	{
		backends.GET("", backendHandler.ListBackends)
		backends.POST("", backendHandler.CreateBackend)
		backends.GET("/:backendId", backendHandler.GetBackend)
		backends.PUT("/:backendId", backendHandler.UpdateBackend)
		backends.DELETE("/:backendId", backendHandler.DeleteBackend)

		// Backend state routes
		backends.GET("/:backendId/state", backendHandler.GetBackendState)
		backends.POST("/:backendId/state", backendHandler.UpdateBackendState)
		backends.DELETE("/:backendId/state", backendHandler.ResetBackendState)
	}

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}