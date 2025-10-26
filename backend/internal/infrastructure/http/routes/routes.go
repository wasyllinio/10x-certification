package routes

import (
	"10x-certification/internal/application"
	"10x-certification/internal/infrastructure/http/middleware"

	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// SetupRoutes sets up all HTTP routes
func SetupRoutes(container *application.Container) *gin.Engine {
	router := gin.New()

	// Middlewares
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC3339Nano),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
	}))
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.CORSMiddleware())

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
