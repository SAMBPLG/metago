package metago

import (
	"fmt"
	"net/http"
	"testing"
)

func TestInvalidCredentialError_Is(t *testing.T) {
	type fields struct {
		BaseError BaseError
	}
	type args struct {
		target error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "ShouldEqual",
			fields: fields{
				BaseError: BaseError{},
			},
			args: args{
				target: ErrInvalidCredential,
			},
			want: true,
		},
		{
			name: "ShoulNotdEqual",
			fields: fields{
				BaseError: BaseError{},
			},
			args: args{
				target: ErrBadRequest,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &InvalidCredentialError{
				BaseError: tt.fields.BaseError,
			}
			if got := e.Is(tt.args.target); got != tt.want {
				t.Errorf("InvalidCredentialError.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBadRequestError_Is(t *testing.T) {
	type fields struct {
		BaseError BaseError
	}
	type args struct {
		target error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "ShouldEqual",
			fields: fields{
				BaseError: BaseError{},
			},
			args: args{
				target: ErrBadRequest,
			},
			want: true,
		},
		{
			name: "ShoulNotdEqual",
			fields: fields{
				BaseError: BaseError{},
			},
			args: args{
				target: ErrInvalidCredential,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &BadRequestError{
				BaseError: tt.fields.BaseError,
			}
			if got := e.Is(tt.args.target); got != tt.want {
				t.Errorf("BadRequestError.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseError_Error(t *testing.T) {
	type fields struct {
		Code    int
		Message string
		Err     error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ShouldEqual",
			fields: fields{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			},
			want: "401 Unauthorized",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := BaseError{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
				Err:     tt.fields.Err,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("BaseError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIError_Error(t *testing.T) {
	type fields struct {
		Code    int
		Message string
		Type    APIErrType
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ShouldEqual",
			fields: fields{
				Code:    http.StatusBadRequest,
				Message: "Bad Request",
				Type:    APIErrTypeUnknown,
			},
			want: "400 Bad Request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiError := APIError{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
				Type:    tt.fields.Type,
			}
			if got := apiError.Error(); got != tt.want {
				t.Errorf("APIError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetabaseErr_Error(t *testing.T) {
	type fields struct {
		Errors         map[string]string
		SpecificErrors map[string][]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ShouldEqual",
			fields: fields{
				Errors: map[string]string{
					"model_id": "model_id is required",
				},
				SpecificErrors: map[string][]string{
					"model_id": {
						"model_id should not be empty and is a string",
					},
				},
			},
			want: "model_id: model_id should not be empty and is a string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := MetabaseErr{
				Errors:         tt.fields.Errors,
				SpecificErrors: tt.fields.SpecificErrors,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("MetabaseErr.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetabaseErr_NotEmpty(t *testing.T) {
	type fields struct {
		Errors         map[string]string
		SpecificErrors map[string][]string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ShouldNotEmpty",
			fields: fields{
				Errors: map[string]string{
					"model_id": "model_id is required",
				},
			},
			want: true,
		},
		{
			name: "ShouldEmpty",
			fields: fields{
				Errors: map[string]string{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := MetabaseErr{
				Errors:         tt.fields.Errors,
				SpecificErrors: tt.fields.SpecificErrors,
			}
			if got := e.NotEmpty(); got != tt.want {
				t.Errorf("MetabaseErr.NotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInvalidCredentialError(t *testing.T) {
	type args struct {
		message string
		err     []error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ShouldThrowError",
			args: args{
				message: "invalid credentials",
				err:     []error{fmt.Errorf("invalid credentials")},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewInvalidCredentialError(tt.args.message, tt.args.err...); (err != nil) != tt.wantErr {
				t.Errorf("NewInvalidCredentialError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewBadRequestError(t *testing.T) {
	type args struct {
		message string
		err     []error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ShouldThrowError",
			args: args{
				message: "bad request",
				err:     []error{fmt.Errorf("bad request")},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewBadRequestError(tt.args.message, tt.args.err...); (err != nil) != tt.wantErr {
				t.Errorf("NewBadRequestError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
