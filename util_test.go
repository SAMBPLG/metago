package metago

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
)

func Test_checkForError(t *testing.T) {
	type args struct {
		resp       *resty.Response
		err        error
		errMessage string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ShouldReturnError",
			args: args{
				resp:       &resty.Response{RawResponse: &http.Response{StatusCode: http.StatusBadGateway}},
				err:        fmt.Errorf("bad gateway"),
				errMessage: "error bad gateway",
			},
			wantErr: true,
		},
		{
			name: "ShouldReturnError",
			args: args{
				resp: &resty.Response{
					RawResponse: &http.Response{StatusCode: http.StatusBadRequest}, Request: resty.New().R()},
				errMessage: "error bad request",
			},
			wantErr: true,
		},
		{
			name: "ShouldNotReturnError",
			args: args{
				resp: &resty.Response{
					RawResponse: &http.Response{StatusCode: http.StatusOK}, Request: resty.New().R()},
				errMessage: "error bad request",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkForError(tt.args.resp, tt.args.err, tt.args.errMessage); (err != nil) != tt.wantErr {
				t.Errorf("checkForError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseAPIErrType(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want APIErrType
	}{
		{
			name: "ShouldEqual",
			args: args{
				err: APIError{},
			},
			want: APIErrTypeUnknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseAPIErrType(tt.args.err); got != tt.want {
				t.Errorf("ParseAPIErrType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinPath(t *testing.T) {
	type args struct {
		separator string
		path      []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ShouldEqual",
			args: args{
				separator: "/",
				path:      []string{"url", "path"},
			},
			want: "url/path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinPath(tt.args.separator, tt.args.path...); got != tt.want {
				t.Errorf("JoinPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
