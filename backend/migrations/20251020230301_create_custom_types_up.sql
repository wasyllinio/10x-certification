-- Migration: Create Custom ENUM Types
-- Created: 2025-10-20 23:03:01
-- Description: Create custom ENUM types for connector types, standards, and audit operations

-- Create connector type enum
CREATE TYPE connector_type AS ENUM ('CCS', 'Type2', 'Chademo');

-- Create connector standard enum
CREATE TYPE connector_standard AS ENUM ('AC_1P', 'AC_3P', 'DC');

-- Create audit operation enum
CREATE TYPE audit_operation AS ENUM ('INSERT', 'UPDATE', 'DELETE');
