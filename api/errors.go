package main

import "fmt"

type ApiError struct {
	Code    int
	Message string
}

func (e *ApiError) Error() string {
	return e.Message
}

func MakeApiError(code int, message string) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
	}
}

type ApiValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ApiValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func MakeApiValidationError(field, message string) *ApiValidationError {
	return &ApiValidationError{
		Field:   field,
		Message: message,
	}
}
