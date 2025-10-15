# Plan bazy danych PostgreSQL - EV Chargers Management System

## 1. Przegląd schematu

Schemat bazy danych został zaprojektowany dla systemu zarządzania stacjami ładowania pojazdów elektrycznych (EV Chargers Management System). System obsługuje zarządzanie fizycznymi stacjami ładowania, ich portami, lokalizacjami oraz automatyczne generowanie punktów ładowania (EVSE).

### Główne założenia:
- Soft delete dla wszystkich encji (pole `deleted_at`)
- Optymistic locking dla zapobiegania race conditions
- Pełny audyt operacji CUD (Create, Update, Delete)
- System autoryzacji z rolami Admin i Owner
- Wydajność <200ms dla 95% requestów
- Single-tenant architecture z możliwością rozbudowy na multi-tenant

## 2. Tabele i ich struktura

### 2.1 Tabela `users` (Użytkownicy)

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    password_salt VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'owner')),
    authorization_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);
```

**Opis pól:**
- `id`: Unikalny identyfikator użytkownika (UUID)
- `email`: Adres email użytkownika (unikalny)
- `password_hash`: Zahashowane hasło (argon2)
- `password_salt`: Sól do haszowania hasła
- `role`: Rola użytkownika (admin/owner)
- `authorization_id`: UUID do mechanizmu "wyloguj ze wszystkich urządzeń"
- `created_at`, `updated_at`: Znaczniki czasowe
- `deleted_at`: Soft delete timestamp

### 2.2 Tabela `chargers` (Stacje ładowania)

```sql
CREATE TABLE chargers (
    id UUID PRIMARY KEY,
    vendor VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    serial_number VARCHAR(255) NOT NULL,
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    location_id UUID NULL REFERENCES locations(id) ON DELETE SET NULL,
    assigned_to_location_at TIMESTAMPTZ NULL,
    last_status_change_at TIMESTAMPTZ NULL,
    version INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);
```

**Opis pól:**
- `id`: Unikalny identyfikator stacji (UUID)
- `vendor`: Producent stacji
- `model`: Model stacji
- `serial_number`: Numer seryjny stacji
- `owner_id`: Właściciel stacji (FK do users)
- `location_id`: Lokalizacja stacji (nullable dla stanu "w magazynie")
- `assigned_to_location_at`: Data przypisania do lokalizacji
- `last_status_change_at`: Data ostatniej zmiany statusu
- `version`: Wersja dla optymistic locking
- `created_at`, `updated_at`: Znaczniki czasowe
- `deleted_at`: Soft delete timestamp

### 2.3 Tabela `locations` (Lokalizacje)

```sql
CREATE TABLE locations (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    country_code CHAR(3) NOT NULL CHECK (country_code ~ '^[A-Z]{3}$'),
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    version INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);
```

**Opis pól:**
- `id`: Unikalny identyfikator lokalizacji (UUID)
- `name`: Nazwa lokalizacji
- `address`: Adres lokalizacji
- `country_code`: Kod kraju (ISO 3166-1 alpha-3)
- `owner_id`: Właściciel lokalizacji (FK do users)
- `version`: Wersja dla optymistic locking
- `created_at`, `updated_at`: Znaczniki czasowe
- `deleted_at`: Soft delete timestamp

### 2.4 Tabela `connectors` (Porty stacji)

```sql
CREATE TYPE connector_type AS ENUM ('CCS', 'Type2', 'Chademo');
CREATE TYPE connector_standard AS ENUM ('AC_1P', 'AC_3P', 'DC');

CREATE TABLE connectors (
    id UUID PRIMARY KEY,
    charger_id UUID NOT NULL REFERENCES chargers(id) ON DELETE CASCADE,
    connector_id INTEGER NOT NULL CHECK (connector_id > 0),
    power DECIMAL(9,1) NOT NULL CHECK (power > 0),
    voltage INTEGER NOT NULL CHECK (voltage > 0),
    amperage INTEGER NOT NULL CHECK (amperage > 0),
    connector_type connector_type NOT NULL,
    connector_standard connector_standard NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,
    CONSTRAINT uq_charger_connector UNIQUE (charger_id, connector_id)
);
```

**Opis pól:**
- `id`: Unikalny identyfikator portu (UUID)
- `charger_id`: Stacja ładowania (FK do chargers)
- `connector_id`: ID portu w ramach stacji
- `power`: Moc portu w kW
- `voltage`: Napięcie w V
- `amperage`: Natężenie w A
- `connector_type`: Typ złącza (ENUM)
- `connector_standard`: Standard złącza (ENUM)
- `created_at`, `updated_at`: Znaczniki czasowe
- `deleted_at`: Soft delete timestamp

### 2.5 Tabela `evse` (Punkty ładowania)

```sql
CREATE TABLE evse (
    id UUID PRIMARY KEY,
    evse_id VARCHAR(50) NOT NULL CHECK (evse_id ~ '^[A-Z]{2}\*[A-Z0-9]{3}\*E[A-Z0-9\*]+$'),
    connector_id UUID UNIQUE NOT NULL REFERENCES connectors(id) ON DELETE CASCADE,
    location_id UUID NOT NULL REFERENCES locations(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);
```

**Opis pól:**
- `id`: Unikalny identyfikator punktu ładowania (UUID)
- `evse_id`: ID punktu ładowania w formacie Emi3spec
- `connector_id`: Port stacji (unique FK do connectors)
- `location_id`: Lokalizacja punktu ładowania (FK do locations)
- `created_at`: Znacznik czasowy utworzenia
- `deleted_at`: Soft delete timestamp

### 2.6 Tabela `audit_logs` (Audyt operacji)

```sql
CREATE TYPE audit_operation AS ENUM ('INSERT', 'UPDATE', 'DELETE');

CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    operation audit_operation NOT NULL,
    table_name VARCHAR(50) NOT NULL,
    record_id UUID NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    old_values JSONB NULL,
    new_values JSONB NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
) PARTITION BY RANGE (created_at);
```

**Opis pól:**
- `id`: Unikalny identyfikator wpisu audytu (UUID)
- `operation`: Typ operacji (ENUM: INSERT/UPDATE/DELETE)
- `table_name`: Nazwa tabeli, na której wykonano operację
- `record_id`: ID rekordu, na którym wykonano operację
- `user_id`: Użytkownik wykonujący operację (FK do users)
- `old_values`: Poprzednie wartości (JSONB)
- `new_values`: Nowe wartości (JSONB)
- `created_at`: Znacznik czasowy operacji

## 3. Relacje między tabelami

### 3.1 Diagram relacji

```
users (1) -----> (N) chargers
users (1) -----> (N) locations
chargers (1) -----> (N) connectors
chargers (N) -----> (1) locations (nullable)
connectors (1) -----> (1) evse (unique)
locations (1) -----> (N) evse
users (1) -----> (N) audit_logs
```

### 3.2 Szczegółowe relacje

1. **users → chargers** (1:N)
   - Jeden użytkownik może mieć wiele stacji ładowania
   - `chargers.owner_id` → `users.id`
   - ON DELETE RESTRICT (ochrona przed utratą danych)

2. **users → locations** (1:N)
   - Jeden użytkownik może mieć wiele lokalizacji
   - `locations.owner_id` → `users.id`
   - ON DELETE RESTRICT

3. **chargers → connectors** (1:N)
   - Jedna stacja może mieć wiele portów
   - `connectors.charger_id` → `chargers.id`
   - ON DELETE CASCADE (usunięcie stacji usuwa porty)

4. **chargers → locations** (N:1, nullable)
   - Stacja może być przypisana do lokalizacji lub być "w magazynie"
   - `chargers.location_id` → `locations.id`
   - ON DELETE SET NULL (usunięcie lokalizacji zwraca stację do magazynu)

5. **connectors → evse** (1:1, unique)
   - Jeden port generuje jeden punkt ładowania
   - `evse.connector_id` → `connectors.id` (UNIQUE)
   - ON DELETE CASCADE

6. **locations → evse** (1:N)
   - Jedna lokalizacja może mieć wiele punktów ładowania
   - `evse.location_id` → `locations.id`
   - ON DELETE CASCADE

7. **users → audit_logs** (1:N)
   - Jeden użytkownik może mieć wiele wpisów audytu
   - `audit_logs.user_id` → `users.id`
   - ON DELETE RESTRICT

## 4. Indeksy

### 4.1 Indeksy unikalne z soft delete

```sql
-- Unikalność email z uwzględnieniem soft delete
CREATE UNIQUE INDEX idx_users_email ON users(email) WHERE deleted_at IS NULL;

-- Unikalność numeru seryjnego w ramach producenta
CREATE UNIQUE INDEX idx_chargers_serial ON chargers(vendor, serial_number) WHERE deleted_at IS NULL;

-- Unikalność connector_id w ramach stacji
CREATE UNIQUE INDEX idx_connectors_unique ON connectors(charger_id, connector_id) WHERE deleted_at IS NULL;

-- Unikalność evse_id
CREATE UNIQUE INDEX idx_evse_id_unique ON evse(evse_id) WHERE deleted_at IS NULL;
```

### 4.2 Indeksy wydajnościowe

```sql
-- Indeksy dla relacji FK
CREATE INDEX idx_chargers_owner ON chargers(owner_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_chargers_location ON chargers(location_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_locations_owner ON locations(owner_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_connectors_charger ON connectors(charger_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evse_connector ON evse(connector_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evse_location ON evse(location_id) WHERE deleted_at IS NULL;

-- Indeksy kompozytowe dla częstych zapytań
CREATE INDEX idx_chargers_owner_location ON chargers(owner_id, location_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_chargers_pagination ON chargers(created_at DESC, id) WHERE deleted_at IS NULL;
CREATE INDEX idx_locations_pagination ON locations(created_at DESC, id) WHERE deleted_at IS NULL;
CREATE INDEX idx_connectors_pagination ON connectors(created_at DESC, id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evse_pagination ON evse(created_at DESC, id) WHERE deleted_at IS NULL;
```

### 4.3 Indeksy dla wyszukiwania

```sql
-- Extension dla full-text search
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- GIN indeks dla wyszukiwania stacji
CREATE INDEX idx_chargers_search ON chargers USING GIN (vendor gin_trgm_ops, model gin_trgm_ops, serial_number gin_trgm_ops);

-- GIN indeks dla wyszukiwania lokalizacji
CREATE INDEX idx_locations_search ON locations USING GIN (name gin_trgm_ops, address gin_trgm_ops);
```

### 4.4 Indeksy dla audytu

```sql
-- Indeksy kompozytowe dla audytu
CREATE INDEX idx_audit_user_time ON audit_logs(user_id, created_at DESC);
CREATE INDEX idx_audit_table_record ON audit_logs(table_name, record_id, created_at DESC);
CREATE INDEX idx_audit_operation_time ON audit_logs(operation, created_at DESC);

-- GIN indeks dla JSONB
CREATE INDEX idx_audit_jsonb ON audit_logs USING GIN (old_values, new_values);
```

## 5. Ograniczenia i walidacje

### 5.1 CHECK constraints

```sql
-- Walidacja ról użytkowników
ALTER TABLE users ADD CONSTRAINT chk_users_role CHECK (role IN ('admin', 'owner'));

-- Walidacja kodu kraju
ALTER TABLE locations ADD CONSTRAINT chk_country_code CHECK (country_code ~ '^[A-Z]{3}$');

-- Walidacja wartości dodatnich dla portów
ALTER TABLE connectors ADD CONSTRAINT chk_connector_id CHECK (connector_id > 0);
ALTER TABLE connectors ADD CONSTRAINT chk_power CHECK (power > 0);
ALTER TABLE connectors ADD CONSTRAINT chk_voltage CHECK (voltage > 0);
ALTER TABLE connectors ADD CONSTRAINT chk_amperage CHECK (amperage > 0);

-- Walidacja formatu EvseID
ALTER TABLE evse ADD CONSTRAINT chk_evse_id_format CHECK (evse_id ~ '^[A-Z]{3}\*[A-Z0-9]+\*E[A-Z0-9\*]+\*[0-9]+$');
```

### 5.2 UNIQUE constraints

```sql
-- Unikalność email (z soft delete)
-- Zdefiniowane w indeksach

-- Unikalność numeru seryjnego w ramach producenta
-- Zdefiniowane w indeksach

-- Unikalność connector_id w ramach stacji
-- Zdefiniowane w tabeli connectors

-- Unikalność evse_id
-- Zdefiniowane w indeksach
```

## 6. Konfiguracja bazy danych

### 6.1 Ustawienia PostgreSQL

```sql
-- Ustawienie strefy czasowej na UTC
SET timezone = 'UTC';

-- Włączenie rozszerzeń
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pg_trgm;
```

### 6.2 Polityki bezpieczeństwa

```sql
-- Row Level Security (opcjonalne - może być implementowane aplikacyjnie)
-- Przykład dla tabeli chargers:
ALTER TABLE chargers ENABLE ROW LEVEL SECURITY;

CREATE POLICY chargers_owner_policy ON chargers
    FOR ALL TO application_role
    USING (
        owner_id = current_setting('app.current_user_id')::UUID
        OR EXISTS (
            SELECT 1 FROM users 
            WHERE id = current_setting('app.current_user_id')::UUID 
            AND role = 'admin'
        )
    );
```

## 7. Uwagi implementacyjne

### 7.1 Soft Delete
- Wszystkie tabele używają pola `deleted_at` dla soft delete
- Indeksy partial z `WHERE deleted_at IS NULL` zapewniają wydajność
- Kaskadowe soft delete implementowane aplikacyjnie

### 7.2 Optymistic Locking
- Pola `version` w tabelach `chargers` i `locations`
- Sprawdzanie wersji przed UPDATE w aplikacji
- Increment wersji przy każdej aktualizacji

### 7.3 Generowanie EvseID
- Format zgodny ze specyfikacją Emi3spec
- Generowanie aplikacyjne przy przypisaniu stacji do lokalizacji
- Walidacja formatu przez CHECK constraint

### 7.4 Audyt operacji
- Logowanie tylko operacji CUD (bez SELECT)
- Przechowywanie old/new values w JSONB
- Partycjonowanie miesięczne dla długoterminowego storage

### 7.5 Wydajność
- Cursor-based pagination z indeksem `(created_at DESC, id)`
- GIN indeksy dla full-text search
- Partial indeksy dla soft delete
- Brak materialized views w MVP

## 8. Migracje

### 8.1 Kolejność migracji
1. Extensions (uuid-ossp, pg_trgm)
2. Custom types (connector_type, connector_standard, audit_operation)
3. Tabele w kolejności dependencies:
   - users
   - chargers
   - locations
   - connectors
   - evse
   - audit_logs
4. Indeksy (partial, GIN, B-tree, compound)
5. Partycjonowanie audit_logs
6. Polityki RLS (opcjonalne)

### 8.2 Format migracji
- Każda migracja w osobnym pliku z timestampem
- Naming: `YYYYMMDDHHMMSS_description.sql`
- Rollback scripts dla każdej migracji
- Testy integracyjne po każdej migracji

Ten schemat bazy danych zapewnia solidną podstawę dla systemu zarządzania stacjami ładowania pojazdów elektrycznych, spełniając wszystkie wymagania funkcjonalne i niefunkcjonalne określone w PRD oraz notatkach z sesji planowania.
