package httpResponses

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Error   string      `json:"error"`
	Status  int         `json:"code,omitempty"`
	Details interface{} `json:"details,omitempty"`
}
