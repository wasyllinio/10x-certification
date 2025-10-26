package postgres

import (
	"10x-certification/internal/config"
	"10x-certification/internal/infrastructure/logger"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection represents PostgreSQL database connection
type Connection struct {
	db     *gorm.DB
	logger *slog.Logger
}

// Connect establishes connection to PostgreSQL database
func Connect(cfg *config.Config) *Connection {
	// Set global logger level based on config
	log := logger.Default().With("source", "postgres")
	log.Info("Connecting to PostgreSQL database")

	// Configure GORM logger
	var gormLogLevel slog.Level
	switch cfg.LogLevel {
	case "debug":
		gormLogLevel = slog.LevelInfo
	case "error":
		gormLogLevel = slog.LevelError
	default:
		gormLogLevel = slog.LevelWarn
	}

	s := slogGorm.New(slogGorm.WithHandler(log.Handler()),
		slogGorm.SetLogLevel(slogGorm.ErrorLogType, gormLogLevel),
		slogGorm.SetLogLevel(slogGorm.SlowQueryLogType, gormLogLevel),
		slogGorm.SetLogLevel(slogGorm.DefaultLogType, gormLogLevel),
	)

	// Configure GORM with PostgreSQL driver
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
		Logger: s,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	// Get underlying sql.DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("Failed to get underlying sql.DB", "error", err)
		os.Exit(1)
	}

	// Configure connection pool settings
	sqlDB.SetMaxOpenConns(25)                 // Maximum number of open connections
	sqlDB.SetMaxIdleConns(10)                 // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Maximum connection lifetime
	sqlDB.SetConnMaxIdleTime(1 * time.Minute) // Maximum idle connection time

	logger.Info("Successfully connected to PostgreSQL database")

	// Get the directory where this file is located and construct path to migrations
	migrationsDir := "file://migrations"

	m, err := migrate.New(
		migrationsDir,
		cfg.DatabaseURL,
	)
	if err != nil {
		logger.Error("Failed to create migration instance", "error", err)
		os.Exit(1)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Error("Failed to run migrations", "error", err)
		os.Exit(1)
	}

	logger.Info("Migrations completed successfully")

	return &Connection{
		db:     db,
		logger: log,
	}
}

// Close closes the database connection
func (c *Connection) Close() error {
	c.logger.Info("Closing database connection...")

	if c.db != nil {
		sqlDB, err := c.db.DB()
		if err != nil {
			c.logger.Error("Failed to get underlying sql.DB for closing", "error", err)
			return err
		}

		if err := sqlDB.Close(); err != nil {
			c.logger.Error("Failed to close database connection", "error", err)
			return err
		}
	}

	c.logger.Info("Database connection closed successfully")
	return nil
}

// Ping checks if database connection is alive
func (c *Connection) Ping(ctx context.Context) error {
	if c.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	sqlDB, err := c.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	return nil
}

// DB returns the GORM database instance
func (c *Connection) DB() *gorm.DB {
	return c.db
}
