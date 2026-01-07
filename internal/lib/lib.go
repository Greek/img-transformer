package lib

import (
	"encoding/json"
	"net/http"
)

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
