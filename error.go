package metago

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrInvalidCredential = &InvalidCredentialError{}
	ErrBadRequest        = &BadRequestError{}
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
	return fmt.Sprintf("%v %v", e.Code, e.Message)
}

const (
	// APIErrTypeUnknown is for API errors that are not strongly
	// typed.
	APIErrTypeUnknown APIErrType = "unknown"
)

// APIError holds message and statusCode for api errors
type APIError struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Type    APIErrType `json:"type"`
}

// Error stringifies the APIError
func (apiError APIError) Error() string {
	return apiError.Message
}

// http error response model
type HTTPErrorResponse struct {
	Error       string `json:"error,omitempty"`
	Message     string `json:"error_message,omitempty"`
	Description string `json:"error_description,omitempty"`
}

// string representation of http error response model
func (err HTTPErrorResponse) String() string {
	var res strings.Builder
	if len(err.Error) > 0 {
		res.WriteString(err.Error)
	}
	if len(err.Message) > 0 {
		if res.Len() > 0 {
			res.WriteString(": ")
		}
		res.WriteString(err.Message)
	}
	if len(err.Description) > 0 {
		if res.Len() > 0 {
			res.WriteString(": ")
		}

		res.WriteString(err.Description)
	}
	return res.String()
}

// NotEmpty validates that error is not empty
func (err HTTPErrorResponse) NotEmpty() bool {
	return len(err.Error) > 0 || len(err.Message) > 0 || len(err.Description) > 0
}

//	{
//		"specific-errors": {
//		   "authority_level": [
//			  "should be \"official\", received: \"\""
//		   ],
//		   "description": [
//			  "should be at least 1 character, received: \"\"",
//			  "non-blank string, received: \"\""
//		   ],
//		   "namespace": [
//			  "should be at least 1 character, received: \"\"",
//			  "non-blank string, received: \"\""
//		   ]
//		},
//		"errors": {
//		   "description": "nullable value must be a non-blank string.",
//		   "namespace": "nullable value must be a non-blank string.",
//		   "authority_level": "nullable enum of official"
//		}
//	 }
type MetabaseErr struct {
	Errors         map[string]string   `json:"errors"`
	SpecificErrors map[string][]string `json:"specific-errors,omitempty"`
}

func (e MetabaseErr) Error() string {
	return ""
}

// []byte representation of error response model
func (err MetabaseErr) JSON() []byte {
	if len(err.Errors) > 0 {
		js, _ := json.Marshal(err.Errors)
		return js
	}
	return make([]byte, 0)
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
