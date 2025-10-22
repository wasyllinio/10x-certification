package routes

import (
	"10x-certification/internal/application"
	"10x-certification/internal/infrastructure/http/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all HTTP routes
func SetupRoutes(container *application.Container) *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(middleware.ErrorMiddleware())

	// Health check
	router.GET("/health", container.HealthHandler.Health)

	// Public routes
	auth := router.Group("/auth")
	{
		auth.POST("/register", container.AuthHandler.Register)
		auth.POST("/login", container.AuthHandler.Login)
	}

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware(container.JWTService))
	{
		// Chargers
		chargers := api.Group("/chargers")
		{
			chargers.GET("", container.ChargerHandler.List)
			chargers.POST("", container.ChargerHandler.Create)
			chargers.GET("/:id", container.ChargerHandler.GetByID)
			chargers.PUT("/:id", container.ChargerHandler.Update)
			chargers.DELETE("/:id", container.ChargerHandler.Delete)

			// Connectors
			chargers.POST("/:id/connectors", container.ChargerHandler.AddConnector)
			chargers.PUT("/connectors/:connectorId", container.ChargerHandler.UpdateConnector)
			chargers.DELETE("/connectors/:connectorId", container.ChargerHandler.DeleteConnector)
		}

		// Locations
		locations := api.Group("/locations")
		{
			locations.GET("", container.LocationHandler.List)
			locations.POST("", container.LocationHandler.Create)
			locations.GET("/:id", container.LocationHandler.GetByID)
			locations.PUT("/:id", container.LocationHandler.Update)
			locations.DELETE("/:id", container.LocationHandler.Delete)

			// Charger assignment
			locations.PUT("/:id/assign-charger", container.LocationHandler.AssignCharger)
			locations.DELETE("/:id/chargers/:chargerId", container.LocationHandler.DetachCharger)
			locations.GET("/:id/chargers", container.LocationHandler.GetChargers)
			locations.GET("/:id/evse", container.LocationHandler.GetEVSE)
		}
	}

	return router
}
