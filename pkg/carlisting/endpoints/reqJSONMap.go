package endpoints

// type GetRequest struct {
// 	Filters []internal.Filter `json:"filters,omitempty"`
// }

type GetResponse struct {
	Int int    `json:"Int"`
	Err string `json:"err,omitempty"`
}

type ListNewCarRequest struct {
	B int `json:"B"`
}

type ListNewCarResponse struct {
	Int int    `json:"Int"`
	Err string `json:"err,omitempty"`
}

type SetAvailabilityRequest struct {
	Dates string `json:"dates"`
}

type SetAvailabilityResponse struct {
	Int int    `json:"Int"`
	Err string `json:"err,omitempty"`
}
