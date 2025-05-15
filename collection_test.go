package metago

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

const apikey string = "yourapikey"

func TestCollectionParameters_Values(t *testing.T) {
	type fields struct {
		Archived                    bool
		ExcludeOtherUserCollections bool
		Namespace                   string
		PersonalOnly                bool
	}
	tests := []struct {
		name   string
		fields fields
		want   url.Values
	}{
		{
			name: "ShouldEqual",
			fields: fields{
				Archived:                    true,
				ExcludeOtherUserCollections: false,
				Namespace:                   "",
				PersonalOnly:                true,
			},
			want: url.Values{"archived": []string{"true"}, "exclude-other-user-collections": []string{"false"}, "personal-only": []string{"true"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CollectionParameters{
				Archived:                    tt.fields.Archived,
				ExcludeOtherUserCollections: tt.fields.ExcludeOtherUserCollections,
				Namespace:                   tt.fields.Namespace,
				PersonalOnly:                tt.fields.PersonalOnly,
			}
			if got := c.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollectionParameters.Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_Collections(t *testing.T) {
	type fields struct {
		sdk        *Metago
		handleFunc handleFunc
	}
	type args struct {
		ctx  context.Context
		opts CollectionParameters
	}
	ctx := context.Background()
	sdk, err := New(Option{
		BasePath:   "http://localhost:3000",
		AuthMethod: APIKEY,
		APIKey:     apikey,
		Context:    ctx,
	})
	if err != nil {
		t.Errorf("Collection.Collections() error = %v", err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ShouldReturnCollections",
			fields: fields{
				sdk: sdk,
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: http.StatusOK}, nil
				},
			},
			args: args{
				ctx:  ctx,
				opts: CollectionParameters{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &collection{
				sdk: tt.fields.sdk,
			}
			m.sdk.restyClient.SetTransport(&RoundTripper{
				HandleFunc: tt.fields.handleFunc,
			})
			var res []Collection
			if err := m.Collections(tt.args.ctx, tt.args.opts, &res); (err != nil) != tt.wantErr {
				t.Errorf("Collection.Collections() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_collection_Root(t *testing.T) {
	type fields struct {
		sdk        *Metago
		handleFunc handleFunc
	}
	type args struct {
		ctx    context.Context
		opts   CollectionParameters
		result *Collection
	}
	ctx := context.Background()
	sdk, err := New(Option{
		BasePath:   "http://localhost:3000",
		AuthMethod: APIKEY,
		APIKey:     apikey,
		Context:    ctx,
	})
	if err != nil {
		t.Errorf("Collection.Collections() error = %v", err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ShouldReturnRootCollection",
			fields: fields{
				sdk: sdk,
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: http.StatusOK}, nil
				},
			},
			args: args{
				ctx:    ctx,
				opts:   CollectionParameters{},
				result: &Collection{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &collection{
				sdk: tt.fields.sdk,
			}
			m.sdk.restyClient.SetTransport(&RoundTripper{HandleFunc: tt.fields.handleFunc})
			if err := m.Root(tt.args.ctx, tt.args.opts, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("collection.Root() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("[%v] %s", tt.args.result.ID, tt.args.result.Name)
		})
	}
}

func Test_collection_Tree(t *testing.T) {
	type fields struct {
		sdk        *Metago
		handleFunc handleFunc
	}
	type args struct {
		ctx    context.Context
		opts   TreeParameters
		result *[]Collection
	}
	ctx := context.Background()
	sdk, err := New(Option{
		BasePath:   "http://localhost:3000",
		AuthMethod: APIKEY,
		APIKey:     apikey,
		Context:    ctx,
	})
	if err != nil {
		t.Errorf("Collection.Collections() error = %v", err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ShouldReturnTree",
			fields: fields{
				sdk: sdk,
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: http.StatusOK}, nil
				},
			},
			args: args{
				ctx:    ctx,
				opts:   TreeParameters{},
				result: &[]Collection{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &collection{
				sdk: tt.fields.sdk,
			}
			m.sdk.restyClient.SetTransport(&RoundTripper{HandleFunc: tt.fields.handleFunc})
			if err := m.Tree(tt.args.ctx, tt.args.opts, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("collection.Tree() error = %v, wantErr %v", err, tt.wantErr)
			}

			// for _, collection := range *tt.args.result {
			// 	// t.Log(collection)
			// 	t.Logf("[%v] %v, Here: %v", collection.ID, collection.Name, collection.Here)

			// 	// if collection.Children != nil && len(collection.Children) > 0 {
			// 	if len(collection.Children) > 0 {
			// 		t.Logf("Children: %v", len(collection.Children))
			// 		for _, c0 := range collection.Children {
			// 			t.Logf("-> [%v] %v", c0.ID, c0.Name)
			// 		}
			// 	}
			// }
		})
	}
}

func Test_collection_Get(t *testing.T) {
	type fields struct {
		sdk        *Metago
		handleFunc handleFunc
	}
	type args struct {
		ctx    context.Context
		id     string
		result *Collection
	}
	ctx := context.Background()
	sdk, err := New(Option{
		BasePath:   "http://localhost:3000",
		AuthMethod: APIKEY,
		APIKey:     apikey,
		Context:    ctx,
	})
	if err != nil {
		t.Errorf("Collection.Collections() error = %v", err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ShouldReturnCollection2",
			fields: fields{
				sdk: sdk,
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: http.StatusOK}, nil
				},
			},
			args: args{
				ctx:    ctx,
				id:     "2",
				result: &Collection{},
			},
		},
		{
			name: "ShouldReturnCollection1",
			fields: fields{
				sdk: sdk,
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: http.StatusOK}, nil
				},
			},
			args: args{
				ctx:    ctx,
				id:     "1",
				result: &Collection{},
			},
		},
		{
			name: "ShouldReturnCollection3",
			fields: fields{
				sdk: sdk,
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: http.StatusOK}, nil
				},
			},
			args: args{
				ctx:    ctx,
				id:     "3",
				result: &Collection{},
			},
		},
		{
			name: "ShouldReturnCollection4",
			fields: fields{
				sdk: sdk,
				handleFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: http.StatusOK}, nil
				},
			},
			args: args{
				ctx:    ctx,
				id:     "4",
				result: &Collection{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &collection{
				sdk: tt.fields.sdk,
			}
			m.sdk.restyClient.SetTransport(&RoundTripper{HandleFunc: tt.fields.handleFunc})
			if err := m.Get(tt.args.ctx, tt.args.id, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("collection.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
