package lib

import "fmt"

// ErrResponse is the interface for generating HTTP errors.
type ErrResponse interface {
	error
	ErrCode() int
	ErrHTTPCode() int
	ErrReason() string
}

// HTTPErr is a custom error type for generating coherent HTTP errors.
type HTTPErr struct {
	Code     int
	HTTPCode int
	Reason   string
}

func (e *HTTPErr) Error() string {
	return fmt.Sprintf("error %d: %s", e.HTTPCode, e.Reason)
}

func (e *HTTPErr) ErrCode() int {
	return e.Code
}

func (e *HTTPErr) ErrHTTPCode() int {
	return e.HTTPCode
}

func (e *HTTPErr) ErrReason() string {
	return e.Reason
}
