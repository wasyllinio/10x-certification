-- Migration: Drop PostgreSQL Extensions
-- Created: 2025-10-20 23:03:00
-- Description: Drop PostgreSQL extensions (reverse of 20251020230300_create_extensions_up.sql)

-- Drop trigram extension
DROP EXTENSION IF EXISTS pg_trgm;

-- Drop UUID generation extension
DROP EXTENSION IF EXISTS "uuid-ossp";
