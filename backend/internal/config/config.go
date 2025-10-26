package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config holds application configuration
type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	LogLevel      string `mapstructure:"LOG_LEVEL"`
	AppEnv        string // ustawiane ręcznie, nie z ENV
}

// Load loads configuration using Viper with godotenv fallback
func Load() (*Config, error) {
	// 1. Określ środowisko (z ENV lub domyślne)
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	// 2. Załaduj odpowiedni plik .env na podstawie APP_ENV
	envFile := fmt.Sprintf(".env.%s", appEnv)
	if _, err := os.Stat(envFile); err == nil {
		if err := godotenv.Load(envFile); err != nil {
			log.Printf("Error loading %s: %v", envFile, err)
		} else {
			log.Printf("Loaded configuration from %s", envFile)
		}
	} else {
		// Fallback do .env
		if _, err := os.Stat(".env"); err == nil {
			if err := godotenv.Load(); err != nil {
				log.Printf("Error loading .env: %v", err)
			} else {
				log.Printf("Loaded configuration from .env")
			}
		}
	}

	// 3. Konfiguracja Viper
	v := viper.New()
	v.AutomaticEnv() // Binduj wszystkie ENV'y
	v.SetEnvPrefix("ECMS")
	v.AllowEmptyEnv(false)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	// 4. Ustaw wartości domyślne
	v.SetDefault("SERVER_ADDRESS", ":8080")
	v.SetDefault("LOG_LEVEL", "info")

	// 5. Unmarshal do struktury
	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %w", err)
	}
	cfg.AppEnv = appEnv

	// 6. Walidacja
	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	log.Printf("Configuration loaded successfully for environment: %s", appEnv)
	return cfg, nil
}

// validateConfig waliduje konfigurację
func validateConfig(cfg *Config) error {
	if cfg.AppEnv == "production" {
		// W production wymagamy wszystkich wartości
		if cfg.JWTSecret == "" || cfg.JWTSecret == "your-secret-key-change-in-production" {
			return fmt.Errorf("JWT_SECRET must be set in production")
		}
		if cfg.DatabaseURL == "" {
			return fmt.Errorf("DATABASE_URL must be set in production")
		}
	} else {
		// W development/test tylko ostrzeżenia
		if cfg.JWTSecret == "" || cfg.JWTSecret == "your-secret-key-change-in-production" {
			log.Printf("WARNING: Using default JWT_SECRET (not suitable for production)")
		}
		if cfg.DatabaseURL == "" {
			log.Printf("WARNING: DATABASE_URL is not set")
		}
	}

	validate := validator.New()
	return validate.Struct(cfg)
}
