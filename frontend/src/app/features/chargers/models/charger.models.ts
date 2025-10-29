// Charger-specific types and interfaces
// These types are used throughout the chargers feature

export type Status = "warehouse" | "assigned";

export type ConnectorType = "CCS" | "Type2" | "Chademo";

export type ConnectorStandard = "AC_1P" | "AC_3P" | "DC";

// Re-export common types that might be shared
export type { 
  ChargerListResponse,
  ChargerDetailResponse,
  CreateChargerRequest,
  UpdateChargerRequest,
  ConnectorResponse,
  ConnectorCreateRequest,
  ConnectorUpdateRequest,
  ChargerListItem,
  ChargerListItemConnector
} from '../../../../types';

