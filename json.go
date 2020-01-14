package errors

// JSON structure for error serialization
type jsError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
