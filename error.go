package metago

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	// ErrInvalidCredential represents invalid credential errors
	ErrInvalidCredential = NewInvalidCredentialError("Unauthorized")

	// ErrBadRequest represents bad request error
	ErrBadRequest = NewBadRequestError("Bad Request")

	// APIErrTypeUnknown is for API errors that are not strongly
	// typed.
	APIErrTypeUnknown APIErrType = "unknown"
)

type MetabaseError struct {
	Error map[string]string `json:"errors,omitempty"`
}

type APIErrType string

type BaseError struct {
	Code    int
	Message string
	Err     error
}

type InvalidCredentialError struct {
	BaseError
}

func (e *InvalidCredentialError) Is(target error) bool {
	_, ok := target.(*InvalidCredentialError)
	return ok
}

type BadRequestError struct {
	BaseError
}

func (e *BadRequestError) Is(target error) bool {
	_, ok := target.(*BadRequestError)
	return ok
}

func (e BaseError) Error() string {
	var errMessage strings.Builder
	errMessage.WriteString(fmt.Sprintf("%v %v", e.Code, e.Message))
	if e.Err != nil {
		errMessage.WriteString(fmt.Sprintf(", %v", e.Err))
	}
	return errMessage.String()
}

// APIError holds message and statusCode for api errors
type APIError struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Type    APIErrType `json:"type"`
}

// Error stringifies the APIError
func (apiError APIError) Error() string {
	return fmt.Sprintf("%v %v", apiError.Code, apiError.Message)
}

type MetabaseErr struct {
	Errors         map[string]string   `json:"errors"`
	SpecificErrors map[string][]string `json:"specific-errors,omitempty"`
}

func (e MetabaseErr) Error() string {
	var errMessage strings.Builder
	for k, v := range e.SpecificErrors {
		if len(v) < 1 {
			continue
		}
		errMessage.WriteString(k)
		errMessage.WriteString(": ")
		errMessage.WriteString(v[0])
	}
	return errMessage.String()
}

func (e MetabaseErr) NotEmpty() bool {
	return len(e.Errors) > 0 || len(e.SpecificErrors) > 0
}

func NewInvalidCredentialError(message string, err ...error) error {
	er := &InvalidCredentialError{
		BaseError: BaseError{
			Code:    http.StatusUnauthorized,
			Message: message,
		},
	}
	for _, e := range err {
		er.Err = e
	}
	return er
}

func NewBadRequestError(message string, err ...error) error {
	er := &BadRequestError{
		BaseError: BaseError{
			Code:    http.StatusBadRequest,
			Message: message,
		},
	}
	for _, e := range err {
		er.Err = e
	}
	return er
}
