package api_response

type APIResponse struct {
	Ok   bool `json:"ok"`
	Data any  `json:"data,omitempty"`
}

func New(ok bool, data any) *APIResponse {
	return &APIResponse{Ok: ok, Data: data}
}
