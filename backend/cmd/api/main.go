package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"10x-certification/internal/application"
	"10x-certification/internal/config"
	"10x-certification/internal/infrastructure/http/routes"
)

// main.go - Application Entry Point
//
// Ten plik jest punktem wejścia aplikacji. Odpowiada za:
// 1. Wczytanie konfiguracji
// 2. Inicjalizację DI Container
// 3. Setup HTTP server z routami
// 4. Graceful shutdown
func main() {
	// 1. Wczytanie konfiguracji
	cfg := config.Load()

	// 2. Inicjalizacja DI Container
	container := application.NewContainer(cfg)

	// 3. Setup HTTP server z routami
	router := routes.SetupRoutes(container)

	// 4. Setup HTTP server
	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: router,
	}

	// 5. Start server in goroutine
	go func() {
		log.Printf("Starting server on %s", cfg.ServerAddress)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server:", err)
		}
	}()

	// 6. Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 7. Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	container.DB.Close()

	log.Println("Server exited")
}
