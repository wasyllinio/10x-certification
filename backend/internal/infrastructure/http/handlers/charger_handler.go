package handlers

import (
	"10x-certification/internal/domain/chargers/command"
	"10x-certification/internal/domain/chargers/dto/request"
	"10x-certification/internal/domain/chargers/query"
	"10x-certification/internal/infrastructure/http/context"
	"10x-certification/internal/infrastructure/http/dto"
	"10x-certification/internal/shared/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChargerHandler handles charger HTTP requests
type ChargerHandler struct {
	// Command handlers
	createChargerHandler   *command.CreateChargerHandler
	updateChargerHandler   *command.UpdateChargerHandler
	deleteChargerHandler   *command.DeleteChargerHandler
	addConnectorHandler    *command.AddConnectorHandler
	updateConnectorHandler *command.UpdateConnectorHandler
	deleteConnectorHandler *command.DeleteConnectorHandler

	// Query handlers
	getChargerByIDHandler *query.GetChargerByIDHandler
	listChargersHandler   *query.ListChargersHandler
	searchChargersHandler *query.SearchChargersHandler
}

// NewChargerHandler creates a new ChargerHandler
func NewChargerHandler(
	createChargerHandler *command.CreateChargerHandler,
	updateChargerHandler *command.UpdateChargerHandler,
	deleteChargerHandler *command.DeleteChargerHandler,
	addConnectorHandler *command.AddConnectorHandler,
	updateConnectorHandler *command.UpdateConnectorHandler,
	deleteConnectorHandler *command.DeleteConnectorHandler,
	getChargerByIDHandler *query.GetChargerByIDHandler,
	listChargersHandler *query.ListChargersHandler,
	searchChargersHandler *query.SearchChargersHandler,
) *ChargerHandler {
	return &ChargerHandler{
		createChargerHandler:   createChargerHandler,
		updateChargerHandler:   updateChargerHandler,
		deleteChargerHandler:   deleteChargerHandler,
		addConnectorHandler:    addConnectorHandler,
		updateConnectorHandler: updateConnectorHandler,
		deleteConnectorHandler: deleteConnectorHandler,
		getChargerByIDHandler:  getChargerByIDHandler,
		listChargersHandler:    listChargersHandler,
		searchChargersHandler:  searchChargersHandler,
	}
}

// Create handles charger creation
func (h *ChargerHandler) Create(c *gin.Context) {
	// Extract user ID from context
	userID, err := context.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
			"code":  "UNAUTHORIZED",
		})
		return
	}

	// Bind and validate request
	var req request.CreateChargerRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		validationErrors := dto.FormatValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "validation failed",
			"code":    "VALIDATION_ERROR",
			"details": validationErrors,
		})
		return
	}

	// Create command with DTO
	cmd := command.NewCreateChargerCommand(&req, userID)

	// Execute command
	chargerID, err := h.createChargerHandler.Handle(c.Request.Context(), cmd)
	if err != nil {
		httpErr := errors.MapDomainErrorToHTTP(err)
		c.JSON(httpErr.StatusCode, gin.H{
			"error": httpErr.Message,
			"code":  httpErr.Code,
		})
		return
	}

	// Return response
	c.JSON(http.StatusCreated, gin.H{
		"id": chargerID.String(),
	})
}

// List handles listing chargers
func (h *ChargerHandler) List(c *gin.Context) {
	// TODO: Implement charger listing HTTP handler
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetByID handles getting charger by ID
func (h *ChargerHandler) GetByID(c *gin.Context) {
	// TODO: Implement get charger by ID HTTP handler
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// Update handles charger update
func (h *ChargerHandler) Update(c *gin.Context) {
	// TODO: Implement charger update HTTP handler
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// Delete handles charger deletion
func (h *ChargerHandler) Delete(c *gin.Context) {
	// TODO: Implement charger deletion HTTP handler
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// AddConnector handles adding connector to charger
func (h *ChargerHandler) AddConnector(c *gin.Context) {
	// TODO: Implement add connector HTTP handler
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// UpdateConnector handles connector update
func (h *ChargerHandler) UpdateConnector(c *gin.Context) {
	// TODO: Implement connector update HTTP handler
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// DeleteConnector handles connector deletion
func (h *ChargerHandler) DeleteConnector(c *gin.Context) {
	// TODO: Implement connector deletion HTTP handler
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
