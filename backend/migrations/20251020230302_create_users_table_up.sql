-- Migration: Create Users Table
-- Created: 2025-10-20 23:03:02
-- Description: Create users table with authentication fields and role management

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    password_salt VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'owner')),
    authorization_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);

-- Add comment for documentation
COMMENT ON TABLE users IS 'User accounts with authentication and authorization information';
COMMENT ON COLUMN users.authorization_id IS 'UUID for logout from all devices mechanism';
COMMENT ON COLUMN users.role IS 'User role: admin or owner';
COMMENT ON COLUMN users.deleted_at IS 'Soft delete timestamp - NULL means active user';
