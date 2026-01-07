package lib

import "fmt"

// ErrResponse is the interface for generating HTTP errors.
type ErrResponse interface {
	error

	ErrHTTPCode() int
	ErrReason() string
}

// httpErr is a custom error type for generating coherent HTTP errors.
type httpErr struct {
	Code     int
	HTTPCode int
	Reason   string
}

func NewErrResponse(reason string, httpCode int) ErrResponse {
	return &httpErr{Reason: reason, HTTPCode: httpCode}
}

func (e *httpErr) Error() string {
	return fmt.Sprintf("error %d: %s", e.HTTPCode, e.Reason)
}

func (e *httpErr) ErrHTTPCode() int {
	return e.HTTPCode
}

func (e *httpErr) ErrReason() string {
	return e.Reason
}
