package digiflazz

import "fmt"

// APIError adalah error yang dikembalikan oleh API Digiflazz.
// Ini terjadi ketika HTTP status code adalah 200, tetapi payload JSON
// berisi response code (rc) yang menandakan kegagalan atau status pending.
type APIError struct {
	ResponseCode ResponseCode `json:"rc"`
	Message      string       `json:"message"`
	Status       string       `json:"status"`
}

// Error mengimplementasikan interface error standar.
func (e *APIError) Error() string {
	return fmt.Sprintf("digiflazz api error: %s (rc: %s, status: %s)", e.Message, e.ResponseCode, e.Status)
}