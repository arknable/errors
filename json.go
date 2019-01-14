package errors

// JSON structure for error serialization
type jsError struct {
	Code    uint16 `json:"code"`
	Message string `json:"message"`
}
