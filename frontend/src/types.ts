// ============================================================================
// ENUMS AND TYPE ALIASES
// ============================================================================

export type Status = "warehouse" | "assigned";

export type ConnectorType = "CCS" | "Type2" | "Chademo";

export type ConnectorStandard = "AC_1P" | "AC_3P" | "DC";

// ============================================================================
// COMMON/SHARED TYPES
// ============================================================================

export interface PaginationResponse {
  page: number;
  limit: number;
  total: number;
  has_next: boolean;
}

export interface IdResponse {
  id: string;
}

// ============================================================================
// AUTHENTICATION DTOs
// ============================================================================

export interface RegisterRequest {
  email: string;
  password: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  token: string;
}

// ============================================================================
// CHARGERS DOMAIN DTOs
// ============================================================================

// Request DTOs
export interface ConnectorCreateRequest {
  connector_id: number;
  power: number;
  voltage: number;
  amperage: number;
  connector_type: ConnectorType;
  connector_standard: ConnectorStandard;
}

export interface ConnectorUpdateRequest {
  id?: string;
  connector_id: number;
  power: number;
  voltage: number;
  amperage: number;
  connector_type: ConnectorType;
  connector_standard: ConnectorStandard;
}

export interface CreateChargerRequest {
  vendor: string;
  model: string;
  serial_number: string;
  connectors: ConnectorCreateRequest[];
}

export interface UpdateChargerRequest {
  vendor: string;
  model: string;
  serial_number: string;
  version: number;
  connectors: ConnectorUpdateRequest[];
}

// Response DTOs
export interface ConnectorResponse {
  id: string;
  connector_id: number;
  power: number;
  voltage: number;
  amperage: number;
  connector_type: ConnectorType;
  connector_standard: ConnectorStandard;
}

export interface ChargerListItemConnector {
  power: number;
  connector_type: ConnectorType;
}

export interface ChargerListItem {
  id: string;
  vendor: string;
  model: string;
  serial_number: string;
  location_id: string | null;
  connectors: ChargerListItemConnector[];
  created_at: Date;
}

export interface ChargerDetailResponse {
  vendor: string;
  model: string;
  serial_number: string;
  location_id: string | null;
  connectors: ConnectorResponse[];
  created_at: Date;
}

export interface ChargerListResponse {
  data: ChargerListItem[];
  pagination: PaginationResponse;
}

// ============================================================================
// LOCATIONS DOMAIN DTOs
// ============================================================================

// Request DTOs
export interface CreateLocationRequest {
  name: string;
  address: string;
  country_code: string;
}

export interface UpdateLocationRequest {
  name: string;
  address: string;
  country_code: string;
  version: number;
}

export interface AssignChargerRequest {
  charger_id: string;
}

// Response DTOs
export interface LocationConnectorResponse {
  id: string;
  connector_id: number;
  power: number;
  voltage: number;
  amperage: number;
  connector_type: ConnectorType;
  connector_standard: ConnectorStandard;
}

export interface LocationChargerResponse {
  id: string;
  vendor: string;
  model: string;
  serial_number: string;
  connectors: LocationConnectorResponse[];
}

export interface EvseResponse {
  id: string;
  evse_id: string;
  connector: LocationConnectorResponse;
  created_at: Date;
}

export interface LocationListItem {
  id: string;
  name: string;
  address: string;
  country_code: string;
  chargers_count: number;
  evse_count: number;
  created_at: Date;
  updated_at: Date;
}

export interface LocationDetailResponse {
  name: string;
  address: string;
  country_code: string;
  version: number;
  chargers: LocationChargerResponse[];
  evses: EvseResponse[];
  created_at: Date;
  updated_at: Date;
}

export interface LocationListResponse {
  data: LocationListItem[];
  pagination: PaginationResponse;
}
