package response

// EVSEListResponse represents a list of EVSEs for a location
type EVSEListResponse struct {
	Data []EVSEResponse `json:"data"`
}

// NewEVSEListResponse creates a new EVSEListResponse
func NewEVSEListResponse(evses []EVSEResponse) *EVSEListResponse {
	return &EVSEListResponse{
		Data: evses,
	}
}
