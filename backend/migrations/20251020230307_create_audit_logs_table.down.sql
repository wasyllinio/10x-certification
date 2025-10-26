-- Migration: Drop Audit Logs Table
-- Created: 2025-10-20 23:03:07
-- Description: Drop audit_logs table (reverse of 20251020230307_create_audit_logs_table_up.sql)

DROP TABLE IF EXISTS audit_logs;
