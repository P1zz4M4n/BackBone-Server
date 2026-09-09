package domain

import (
	"fmt"
)

type Err struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Cause   error  `json:"-"` // Nascosto nelle risposte JSON
}

// Implementation of the error interface
func (e *Err) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

// Unwrap to allow errors.Is and errors.As to work with the underlying cause
func (e *Err) Unwrap() error {
	return e.Cause
}

type ErrCode int

const (
	ErrCodeInternalServer ErrCode = iota + 1
	ErrCodeInvalidInput
	ErrCodeNotFound
	ErrCodeUnauthorized
	ErrCodeForbidden
	ErrCodeConflict
	ErrCodeTooManyRequests
	ErrCodeServiceUnavailable
)

// Automapping from ErrCode to string for the Type field in Err
var errCodeTypeMap = map[ErrCode]string{
	ErrCodeInternalServer:     "INTERNAL_SERVER_ERROR",
	ErrCodeInvalidInput:       "INVALID_INPUT",
	ErrCodeNotFound:           "NOT_FOUND",
	ErrCodeUnauthorized:       "UNAUTHORIZED",
	ErrCodeForbidden:          "FORBIDDEN",
	ErrCodeConflict:           "CONFLICT",
	ErrCodeTooManyRequests:    "TOO_MANY_REQUESTS",
	ErrCodeServiceUnavailable: "SERVICE_UNAVAILABLE",
}

func NewErr(code ErrCode, message string) *Err {
	errType, ok := errCodeTypeMap[code]
	if !ok {
		errType = "UNKNOWN_ERROR"
	}

	return &Err{
		Code:    int(code),
		Type:    errType,
		Message: message,
	}
}

// Specific error constructors for common error types
func NotFoundError(message string) *Err           { return NewErr(ErrCodeNotFound, message) }
func InvalidInputError(message string) *Err       { return NewErr(ErrCodeInvalidInput, message) }
func UnauthorizedError(message string) *Err       { return NewErr(ErrCodeUnauthorized, message) }
func ForbiddenError(message string) *Err          { return NewErr(ErrCodeForbidden, message) }
func ConflictError(message string) *Err           { return NewErr(ErrCodeConflict, message) }
func TooManyRequestsError(message string) *Err    { return NewErr(ErrCodeTooManyRequests, message) }
func ServiceUnavailableError(message string) *Err { return NewErr(ErrCodeServiceUnavailable, message) }

// InternalServerError creates an error representing an internal server error, optionally wrapping a cause.
func InternalServerError(cause error) *Err {
	return &Err{
		Code:    int(ErrCodeInternalServer),
		Type:    errCodeTypeMap[ErrCodeInternalServer],
		Message: "internal server error",
		Cause:   cause,
	}
}
