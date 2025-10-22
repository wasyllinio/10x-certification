# Migrations folder

This folder contains database migration files for PostgreSQL.

## Structure:
- Migration files are named with timestamp: `YYYYMMDDHHMMSS_description.sql`
- Each migration has separate UP and DOWN files for reversibility
- Migrations should be run in order

## Created migration files:

### Extensions and Types
- `20251020230300_create_extensions_up.sql` / `20251020230300_create_extensions_down.sql` - PostgreSQL extensions (uuid-ossp, pg_trgm)
- `20251020230301_create_custom_types_up.sql` / `20251020230301_create_custom_types_down.sql` - Custom ENUM types

### Tables
- `20251020230302_create_users_table_up.sql` / `20251020230302_create_users_table_down.sql` - Users table
- `20251020230303_create_locations_table_up.sql` / `20251020230303_create_locations_table_down.sql` - Locations table
- `20251020230304_create_chargers_table_up.sql` / `20251020230304_create_chargers_table_down.sql` - Chargers table
- `20251020230305_create_connectors_table_up.sql` / `20251020230305_create_connectors_table_down.sql` - Connectors table
- `20251020230306_create_evse_table_up.sql` / `20251020230306_create_evse_table_down.sql` - EVSE table
- `20251020230307_create_audit_logs_table_up.sql` / `20251020230307_create_audit_logs_table_down.sql` - Audit logs table

### Indexes
- `20251020230308_create_performance_indexes_up.sql` / `20251020230308_create_performance_indexes_down.sql` - Performance indexes
- `20251020230309_create_additional_indexes_up.sql` / `20251020230309_create_additional_indexes_down.sql` - Additional composite indexes

## Migration Order:
1. Extensions (uuid-ossp, pg_trgm)
2. Custom types (ENUMs)
3. Tables in dependency order: users → locations → chargers → connectors → evse → audit_logs
4. Performance indexes
5. Additional composite indexes

## Key Features:
- Soft delete support (`deleted_at` column)
- Optimistic locking (`version` column for chargers/locations)
- Full-text search capabilities (GIN indexes)
- Comprehensive audit logging
- Foreign key constraints with proper ON DELETE behavior
