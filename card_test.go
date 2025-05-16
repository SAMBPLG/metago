package metago

import (
	"context"
	"net/http"
	"testing"
)

func TestCardF_String(t *testing.T) {
	tests := []struct {
		name string
		c    CardF
		want string
	}{
		{
			name: "ShouldReturnCorrectString",
			c:    CardFilterArchived,
			want: "archived",
		},
		{
			name: "ShouldReturnCorrectString1",
			c:    CardFilterAll,
			want: "all",
		},
		{
			name: "ShouldReturnCorrectString1",
			c:    CardFilterUsingModel,
			want: "using_model",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("CardF.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardF_EnumIndex(t *testing.T) {
	tests := []struct {
		name string
		c    CardF
		want uint8
	}{
		{
			name: "ShouldReturnCorrectIndex",
			c:    CardFilterAll,
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.EnumIndex(); got != tt.want {
				t.Errorf("CardF.EnumIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_Collections(t *testing.T) {
	type fields struct {
		sdk        *Metago
		handleFunc handleFunc
	}
	type args struct {
		ctx    context.Context
		opts   CardParameters
		result *[]Card
	}
	ctx := context.Background()
	sdk, err := New(Option{
		BasePath:   "http://localhost:3000",
		AuthMethod: APIKEY,
		APIKey:     apikey,
		Context:    ctx,
	})
	if err != nil {
		t.Errorf("card.Collections() error = %v", err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ShouldReturnCardCollection",
			fields: fields{
				sdk: sdk,
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: 201}, nil
				},
			},
			args: args{
				ctx:    ctx,
				opts:   CardParameters{},
				result: &[]Card{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &card{
				sdk: tt.fields.sdk,
			}
			m.sdk.restyClient.SetTransport(&RoundTripper{HandleFunc: tt.fields.handleFunc})
			if err := m.Collections(tt.args.ctx, tt.args.opts, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("card.Collections() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
