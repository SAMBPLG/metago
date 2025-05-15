package metago

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
)

const (
	username string = "your@mail.com"
	password string = "your@password"
)

func TestMetago_Login(t *testing.T) {
	type fields struct {
		basePath    string
		authMethod  AuthMethod
		restyClient *resty.Client
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Metago{
				basePath:    tt.fields.basePath,
				authMethod:  tt.fields.authMethod,
				restyClient: tt.fields.restyClient,
				endpoint:    tt.fields.endpoint,
			}
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
			m.Session = &session{sdk: m}
			// m.restyClient.Debug = true

			ses, err := m.Session.Login(tt.args.ctx, tt.args.usr, tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("Metago.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
			if ses == nil || *ses == "" {
				t.Errorf("Metago.Login() error = %v", fmt.Errorf("session id is empty"))
			}
			m.session = ses
			err = m.Session.Logout(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Metago.Logout() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
