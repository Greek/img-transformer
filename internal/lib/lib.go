package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrResponse is the interface for generating HTTP errors.
type ErrResponse interface {
	error
	ErrCode() int
	ErrReason() string
}

// JSONResponse represents a standard JSON response with data and error fields
type JSONResponse struct {
	Data any `json:"data"`
	Err  any `json:"error"`
}

// WriteJSONSuccess writes a successful JSON response with data
func WriteJSONSuccess(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := JSONResponse{
		Data: data,
		Err:  nil,
	}

	return json.NewEncoder(w).Encode(response)
}

// WriteJSONError writes an error JSON response
func WriteJSONError(w http.ResponseWriter, statusCode int, err any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := JSONResponse{
		Data: nil,
		Err:  err,
	}

	return json.NewEncoder(w).Encode(response)
}

// HTTPErr is a custom error type for generating coherent HTTP errors.
type HTTPErr struct {
	Code   int
	Reason string
}

func (e *HTTPErr) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Reason)
}

func (e *HTTPErr) ErrCode() int {
	return e.Code
}

func (e *HTTPErr) ErrReason() string {
	return e.Reason
}
