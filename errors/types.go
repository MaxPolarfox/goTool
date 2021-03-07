package errors

// Error is the error object for an apiclients request that includes the http
// status code in response
type Error struct {
	Message    string `json:"message"`
}
func (e *Error) Error() string { return e.Message }
