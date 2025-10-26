-- Migration: Create PostgreSQL Extensions
-- Created: 2025-10-20 23:03:00
-- Description: Create required PostgreSQL extensions for UUID generation and full-text search

-- Enable UUID generation extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Enable trigram extension for full-text search
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Set timezone to UTC for consistency
SET timezone = 'UTC';
