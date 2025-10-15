# REST API Plan

## 1. Resources

### Users
- **Database Table**: `users`
- **Description**: System users with admin/owner roles and JWT authentication

### Chargers
- **Database Table**: `chargers`
- **Description**: Physical EV charging stations with vendor, model, and serial number

### Locations
- **Database Table**: `locations`
- **Description**: Physical locations where charging stations can be deployed

### Connectors
- **Database Table**: `connectors`
- **Description**: Charging ports of stations with power, voltage, and connector specifications

### EVSE (Electric Vehicle Supply Equipment)
- **Database Table**: `evse`
- **Description**: Automatically generated charging points from station connectors

### Audit Logs
- **Database Table**: `audit_logs`
- **Description**: System operation audit trail with old/new values

## 2. Endpoints

### Authentication

#### POST /auth/register
- **Description**: Register a new user account
- **Request Body**:
```json
{
  "email": "user@example.com",
  "password": "securepassword123"
}
```
- **Response** (201 Created):
```json
{}
```
- **Error Codes**: 400 (validation error), 409 (email exists)

#### POST /auth/login
- **Description**: Authenticate user and get JWT token
- **Request Body**:
```json
{
  "email": "user@example.com",
  "password": "securepassword123"
}
```
- **Response** (200 OK):
```json
{
  "token": "jwt_token_here"
}
```
- **Error Codes**: 401 (invalid credentials)

### Chargers

#### GET /chargers
- **Description**: List charging stations with pagination and search
- **Query Parameters**:
  - `page` (int): Page number (default: 1)
  - `limit` (int): Items per page (default: 20, max: 100)
  - `search` (string): Search in vendor, model, serial_number
  - `status` (string): Filter by status ("warehouse" or "assigned")
  - `location_id` (uuid): Filter by location
- **Response** (200 OK):
```json
{
  "data": [
    {
      "id": "uuid",
      "vendor": "Tesla",
      "model": "Supercharger V3",
      "serial_number": "TS001",
      "owner_id": "uuid",
      "location_id": "uuid",
      "status": "assigned",
      "connectors": [
        {
          "id": "uuid",
          "connector_id": 1,
          "power": 250.0,
          "voltage": 400,
          "amperage": 625,
          "connector_type": "CCS",
          "connector_standard": "DC"
        }
      ],
      "created_at": "2025-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 100,
    "has_next": true
  }
}
```
- **Error Codes**: 401 (unauthorized)

#### POST /chargers
- **Description**: Create a new charging station
- **Request Body**:
```json
{
  "vendor": "Tesla",
  "model": "Supercharger V3",
  "serial_number": "TS001",
  "connectors": [
    {
      "connector_id": 1,
      "power": 250.0,
      "voltage": 400,
      "amperage": 625,
      "connector_type": "CCS",
      "connector_standard": "DC"
    }
  ]
}
```
- **Response** (200 Created):
```json
{
  "id": "uuid"
}
```
- **Error Codes**: 400 (validation error), 401 (unauthorized), 409 (serial number exists)

#### GET /chargers/{id}
- **Description**: Get charging station details
- **Response** (200 OK):
```json
{
  "id": "uuid",
  "vendor": "Tesla",
  "model": "Supercharger V3",
  "serial_number": "TS001",
  "location_id": "uuid", // nullable
  "connectors": [
    {
      "connector_id": 1,
      "power": 250.0,
      "voltage": 400,
      "amperage": 625,
      "connector_type": "CCS",
      "connector_standard": "DC"
    }
  ]
}
```
- **Error Codes**: 401 (unauthorized), 403 (access denied), 404 (not found)

#### PUT /chargers/{id}
- **Description**: Update charging station (with optimistic locking)
- **Request Body**:
```json
{
  "vendor": "Tesla",
  "model": "Supercharger V4",
  "serial_number": "TS001",
  "version": 1,
  "connectors": [...]
}
```
- **Response** (204 OK): Updated charger object
- **Error Codes**: 400 (validation error), 401 (unauthorized), 403 (access denied), 404 (not found), 409 (version conflict)

#### DELETE /chargers/{id}
- **Description**: Soft delete charging station
- **Response** (204 No Content)
- **Error Codes**: 400 (station assigned to location), 401 (unauthorized), 403 (access denied), 404 (not found)

### Locations

#### GET /locations
- **Description**: List locations with pagination and search
- **Query Parameters**:
  - `page` (int): Page number
  - `limit` (int): Items per page
  - `search` (string): Search in name, address
  - `country_code` (string): Filter by country code
- **Response** (200 OK):
```json
{
  "data": [
    {
      "id": "uuid",
      "name": "Downtown Charging Hub",
      "address": "123 Main St, City",
      "country_code": "USA",
      "owner_id": "uuid",
      "chargers_count": 5,
      "evse_count": 10,
      "created_at": "2025-01-01T00:00:00Z",
      "updated_at": "2025-01-01T00:00:00Z"
    }
  ],
  "pagination": {...}
}
```
- **Error Codes**: 401 (unauthorized)

#### POST /locations
- **Description**: Create a new location
- **Request Body**:
```json
{
  "name": "Downtown Charging Hub",
  "address": "123 Main St, City",
  "country_code": "USA"
}
```
- **Response** (201 Created): Location object
- **Error Codes**: 400 (validation error), 401 (unauthorized)

#### GET /locations/{id}
- **Description**: Get location details
- **Response** (200 OK):
```json
{
  "name": "Updated Name",
  "address": "New Address",
  "country_code": "USA",
  "version": 1,
  "chargers": [...],
  "evses": [...]
}
```
- **Error Codes**: 401 (unauthorized), 403 (access denied), 404 (not found)

#### PUT /locations/{id}
- **Description**: Update location (with optimistic locking)
- **Request Body**:
```json
{
  "name": "Updated Name",
  "address": "New Address",
  "country_code": "USA",
  "version": 1
}
```
- **Response** (200 OK): Updated location
- **Error Codes**: 400 (validation error), 401 (unauthorized), 403 (access denied), 404 (not found), 409 (version conflict)

#### DELETE /locations/{id}
- **Description**: Soft delete location (requires no assigned chargers)
- **Response** (204 No Content)
- **Error Codes**: 400 (has assigned chargers), 401 (unauthorized), 403 (access denied), 404 (not found)

#### PUT /locations/{id}/assign-charger
- **Description**: Assign charger to location (triggers EVSE generation)
- **Request Body**:
```json
{
  "charger_id": "uuid"
}
```
- **Response** (200 OK): Add charger to location_id
- **Error Codes**: 400 (already assigned), 401 (unauthorized), 403 (access denied), 404 (charger/location not found)

#### DELETE /locations/{id}/chargers/{charger_id}
- **Description**: Detach charger from location
- **Response** (204 No Content)
- **Error Codes**: 400 (charger not assigned to the location), 401 (unauthorized), 403 (access denied)

#### GET /locations/{id}/chargers
- **Description**: Get chargers assigned to location
- **Response** (200 OK): Array of charger objects

#### GET /locations/{id}/evse
- **Description**: Get EVSE points for location
- **Response** (200 OK):
```json
{
  "data": [
    {
      "id": "uuid",
      "evse_id": "US*ABC*E123*1",
      "connector": {
        "id": "uuid",
        "connector_id": 1,
        "power": 250.0,
        "voltage": 400,
        "amperage": 625,
        "connector_type": "CCS",
        "connector_standard": "DC"
      },
      "created_at": "2025-01-01T00:00:00Z"
    }
  ]
}
```

## 3. Authentication and Authorization

### JWT Authentication
- **Token Format**: JWT with authorization ID, email, and role
- **Header**: `Authorization: Bearer <token>`
- **Token Expiry**: 24 hours
- **Refresh**: Not implemented in MVP

### Role-Based Authorization
- **Admin Role**: Full access to all resources regardless of ownership
- **Owner Role**: Access only to own resources (chargers, locations, EVSE)
- **Middleware**: Authorization middleware checks JWT and enforces role-based access

### Security Headers
- **CORS**: Configured for Angular frontend
- **Input Validation**: All inputs validated against schema constraints

## 4. Validation and Business Logic

### Validation Rules

#### Users
- Email format validation and uniqueness
- Password minimum 8 characters
- Role must be 'admin' or 'owner'

#### Chargers
- Vendor, model, serial_number required
- Serial number unique per vendor
- Connectors must have positive power, voltage, amperage
- Connector types: CCS, Type2, Chademo
- Connector standards: AC_1P, AC_3P, DC

#### Locations
- Name and address required
- Country code must match ISO 3166-1 alpha-3 format (3 uppercase letters)

#### EVSE
- EvseID must match Emi3spec format: `^[A-Z]{2}\*[A-Z0-9]{3}+\*E[A-Z0-9\*]+$`
- Automatically generated when charger assigned to location

### Business Logic Implementation

#### Automatic EVSE Generation
- Triggered when charger is assigned to location via `PUT /locations/{id}/assign-charger`
- Creates 1:1 mapping between connectors and EVSE points
- Generates EvseID according to Emi3spec format
- Logs generation in audit trail

#### Soft Delete
- All DELETE operations set `deleted_at` timestamp
- Queries filter out soft-deleted records with `WHERE deleted_at IS NULL`
- Cascade soft delete for related entities (connectors, EVSE)

#### Optimistic Locking
- Chargers and locations have `version` field
- PUT requests must include current version
- Version conflict returns 409 status code
- Version incremented on successful update

#### Status Management
- Chargers have status: "warehouse" (location_id = null) or "assigned" (location_id != null)
- Status automatically updated when assigning/unassigning from location
- Cannot delete charger if assigned to location

#### Audit Trail
- All CUD operations logged to audit_logs table
- Stores old and new values in JSONB format
- Includes user ID, operation type, table name, and timestamp
- Partitioned by month for performance

#### Pagination
- Cursor-based pagination using `(created_at DESC, id)` index
- Prevents duplicate results when new records added during pagination
- Default 20 items per page, maximum 100

#### Search and Filtering
- Full-text search using PostgreSQL GIN indexes with trigram matching
- Search across multiple fields (vendor, model, serial_number for chargers)
- Real-time search as user types
- Filters combined with search terms
