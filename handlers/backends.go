package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zij/interstate/service"
)

// BackendHandler handles HTTP requests related to backends
type BackendHandler struct {
	service *service.BackendService
}

// NewBackendHandler creates a new BackendHandler
func NewBackendHandler(service *service.BackendService) *BackendHandler {
	return &BackendHandler{
		service: service,
	}
}

// ListBackends handles GET /api/backends
func (h *BackendHandler) ListBackends(c *gin.Context) {
	backends, err := h.service.ListBackends()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list backends"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"backends": backends})
}

// CreateBackend handles POST /api/backends
func (h *BackendHandler) CreateBackend(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateBackend(input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create backend"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Backend created successfully"})
}

// GetBackend handles GET /api/backends/:backendId
func (h *BackendHandler) GetBackend(c *gin.Context) {
	backendID := c.Param("backendId")
	backend, err := h.service.GetBackend(backendID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Backend not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"backend": backend})
}

// UpdateBackend handles PUT /api/backends/:backendId
func (h *BackendHandler) UpdateBackend(c *gin.Context) {
	backendID := c.Param("backendId")
	var input struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateBackend(backendID, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update backend"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Backend updated successfully"})
}

// DeleteBackend handles DELETE /api/backends/:backendId
func (h *BackendHandler) DeleteBackend(c *gin.Context) {
	backendID := c.Param("backendId")
	err := h.service.DeleteBackend(backendID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete backend"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Backend deleted successfully"})
}

// GetBackendState handles GET /api/backends/:backendId/state
func (h *BackendHandler) GetBackendState(c *gin.Context) {
	backendID := c.Param("backendId")
	state, err := h.service.GetBackendState(backendID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve backend state"})
		return
	}
	c.JSON(http.StatusOK, state)
}

// UpdateBackendState handles POST /api/backends/:backendId/state
func (h *BackendHandler) UpdateBackendState(c *gin.Context) {
	backendID := c.Param("backendId")
	
	var state map[string]interface{}
	if err := c.ShouldBindJSON(&state); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state data"})
		return
	}

	err := h.service.UpdateBackendState(backendID, state)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update backend state"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Backend state updated successfully"})
}

// ResetBackendState handles DELETE /api/backends/:backendId/state
func (h *BackendHandler) ResetBackendState(c *gin.Context) {
	backendID := c.Param("backendId")
	err := h.service.ResetBackendState(backendID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset backend state"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Backend state reset successfully"})
}