package metago

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
)

const (
	username string = "your@mail.com"
	password string = "your@password"
)

type handleFunc = func(req *http.Request) (*http.Response, error)

type RoundTripper struct {
	HandleFunc handleFunc
}

func (r *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.HandleFunc(req)
}

func TestMetago_Login(t *testing.T) {
	type fields struct {
		basePath    string
		authMethod  AuthMethod
		restyClient *resty.Client
		handleFunc  handleFunc
		endpoint    Endpoint
	}
	type args struct {
		ctx  context.Context
		usr  string
		pass string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ShouldNotError",
			fields: fields{
				basePath:    "http://localhost:3000",
				authMethod:  USERPASS,
				restyClient: resty.New(),
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
				endpoint: Endpoint{
					SessionEndpoint: JoinPath("/", "http://localhost:3000", "api", "session"),
				},
			},
			args: args{
				ctx:  context.Background(),
				usr:  username,
				pass: password,
			},
		},
		{
			name: "ShouldErrorInvalidCredentials",
			fields: fields{
				basePath:    "http://localhost:3000",
				authMethod:  USERPASS,
				restyClient: resty.New(),
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return nil, nil
				},
				endpoint: Endpoint{
					SessionEndpoint: JoinPath("/", "http://localhost:3000", "api", "session"),
				},
			},
			args: args{
				ctx:  context.Background(),
				usr:  "address@mail.com",
				pass: "pass",
			},
			wantErr: true,
		},
		{
			name: "ShouldErrorBadRequest",
			fields: fields{
				basePath:    "http://localhost:3000",
				authMethod:  USERPASS,
				restyClient: resty.New(),
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusBadRequest,
					}, nil
				},
				endpoint: Endpoint{
					SessionEndpoint: JoinPath("/", "http://localhost:3000", "api", "session"),
				},
			},
			args: args{
				ctx:  context.Background(),
				usr:  username,
				pass: password,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Metago{
				basePath:    tt.fields.basePath,
				authMethod:  tt.fields.authMethod,
				restyClient: tt.fields.restyClient,
				endpoint:    tt.fields.endpoint,
			}
			m.restyClient.SetTransport(&RoundTripper{
				HandleFunc: tt.fields.handleFunc,
			})
			m.Session = &session{sdk: m}
			_, err := m.Session.Login(tt.args.ctx, tt.args.usr, tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("Metago.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMetago_Logout(t *testing.T) {
	type fields struct {
		basePath    string
		authMethod  AuthMethod
		restyClient *resty.Client
		handleFunc  handleFunc
		endpoint    Endpoint
	}
	type args struct {
		ctx  context.Context
		usr  string
		pass string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ShouldNotError",
			fields: fields{
				basePath:    "http://localhost:3000",
				authMethod:  USERPASS,
				restyClient: resty.New(),
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
				endpoint: Endpoint{
					SessionEndpoint: JoinPath("/", "http://localhost:3000", "api", "session"),
				},
			},
			args: args{
				ctx:  context.Background(),
				usr:  username,
				pass: password,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Metago{
				basePath:    tt.fields.basePath,
				authMethod:  tt.fields.authMethod,
				restyClient: tt.fields.restyClient,
				endpoint:    tt.fields.endpoint,
			}
			sess := ""
			m.session = &sess
			m.Session = &session{sdk: m}
			m.restyClient.SetTransport(&RoundTripper{HandleFunc: tt.fields.handleFunc})

			err := m.Session.Logout(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Metago.Logout() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
