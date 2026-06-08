package utils


type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

// NewError builds a standard error body for a given HTTP status.
func NewError(message string, status int) ErrorResponse {
	return ErrorResponse{Error: message, Status: status}
}