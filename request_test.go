package metago

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestMetago_getRequest(t *testing.T) {
	type fields struct {
		restyClient *resty.Client
	}
	type args struct {
		ctx context.Context
	}
	restclient := resty.New()
	ctx := context.Background()
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *resty.Request
	}{
		{
			name: "ShouldNotError",
			fields: fields{
				restyClient: restclient,
			},
			args: args{
				ctx: ctx,
			},
			want: restclient.R().SetContext(ctx),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Metago{
				restyClient: tt.fields.restyClient,
			}
			if got := m.getRequest(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Metago.getRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetago_getRequestWithAuth(t *testing.T) {
	type fields struct {
		restyClient *resty.Client
		authMethod  AuthMethod
		session     *string
	}
	type args struct {
		ctx context.Context
	}
	restclient := resty.New()
	ctx := context.Background()
	sess := "dummysessionid"
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *resty.Request
	}{
		{
			name: "ShouldNotError_APIKEY",
			fields: fields{
				restyClient: restclient,
				authMethod:  APIKEY,
				session:     &sess,
			},
			args: args{
				ctx: ctx,
			},
			want: restclient.R().SetContext(ctx).SetHeader("x-api-key", sess),
		},
		{
			name: "ShouldNotError_USERPASS",
			fields: fields{
				restyClient: restclient,
				authMethod:  USERPASS,
				session:     &sess,
			},
			args: args{
				ctx: ctx,
			},
			want: restclient.R().SetContext(ctx).SetHeader("x-metabase-session", sess),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Metago{
				restyClient: tt.fields.restyClient,
				authMethod:  tt.fields.authMethod,
				session:     tt.fields.session,
			}
			if got := m.getRequestWithAuth(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Metago.getRequestWithAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
