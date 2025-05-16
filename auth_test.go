package metago

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestMetago_getSession(t *testing.T) {
	type fields struct {
		basePath    string
		authMethod  AuthMethod
		session     *string
		restyClient *resty.Client
		endpoint    Endpoint
		Session     *session
		handleFunc  handleFunc
	}
	type args struct {
		ctx  context.Context
		key  string
		usr  string
		pass string
	}
	resclient := resty.New()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ShouldNotError_APIKEY",
			fields: fields{
				basePath:    "http://localhost:3000",
				authMethod:  APIKEY,
				restyClient: resclient,
				endpoint:    Endpoint{SessionEndpoint: "http://localhost:3000/api/session"},
				Session:     &session{sdk: &Metago{restyClient: resclient, endpoint: Endpoint{SessionEndpoint: "http://localhost:3000/api/session"}}},
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: http.StatusOK}, nil
				},
			},
			args: args{
				ctx: context.Background(),
				key: apikey,
			},
			wantErr: false,
		},
		{
			name: "ShouldNotError_USERPASS",
			fields: fields{
				basePath:    "http://localhost:3000",
				authMethod:  USERPASS,
				restyClient: resclient,
				endpoint:    Endpoint{SessionEndpoint: "http://localhost:3000/api/session"},
				Session:     &session{sdk: &Metago{restyClient: resclient, endpoint: Endpoint{SessionEndpoint: "http://localhost:3000/api/session"}}},
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: http.StatusOK}, nil
				},
			},
			args: args{
				ctx: context.Background(),
				key: apikey,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Metago{
				basePath:    tt.fields.basePath,
				authMethod:  tt.fields.authMethod,
				session:     tt.fields.session,
				restyClient: tt.fields.restyClient,
				endpoint:    tt.fields.endpoint,
				Session:     tt.fields.Session,
			}
			// m.restyClient.Debug = true
			m.restyClient.SetTransport(&RoundTripper{HandleFunc: tt.fields.handleFunc})
			if err := m.getSession(tt.args.ctx, tt.args.key, tt.args.usr, tt.args.pass); (err != nil) != tt.wantErr {
				t.Errorf("Metago.getSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
