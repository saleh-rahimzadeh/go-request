package request

import (
	net_url "net/url"
	"reflect"
	"testing"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

func TestBuildDemand(t *testing.T) {
	type args struct {
		method string
		url    string
		path   string
	}
	tests := []struct {
		name string
		args args
		want Demand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildDemand(tt.args.method, tt.args.url, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildDemand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemand_GetUrl(t *testing.T) {
	type fields struct {
		URI     net_url.URL
		Token   string
		Type    string
		Method  string
		Headers map[string]string
		Error   error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Demand{
				URI:     tt.fields.URI,
				Token:   tt.fields.Token,
				Type:    tt.fields.Type,
				Method:  tt.fields.Method,
				Headers: tt.fields.Headers,
				Error:   tt.fields.Error,
			}
			if got := c.GetUrl(); got != tt.want {
				t.Errorf("Demand.GetUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemand_ContentType(t *testing.T) {
	type fields struct {
		URI     net_url.URL
		Token   string
		Type    string
		Method  string
		Headers map[string]string
		Error   error
	}
	type args struct {
		ctype ContentType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Demand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Demand{
				URI:     tt.fields.URI,
				Token:   tt.fields.Token,
				Type:    tt.fields.Type,
				Method:  tt.fields.Method,
				Headers: tt.fields.Headers,
				Error:   tt.fields.Error,
			}
			if got := c.ContentType(tt.args.ctype); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Demand.ContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemand_AuthorizationBearer(t *testing.T) {
	type fields struct {
		URI     net_url.URL
		Token   string
		Type    string
		Method  string
		Headers map[string]string
		Error   error
	}
	type args struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Demand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Demand{
				URI:     tt.fields.URI,
				Token:   tt.fields.Token,
				Type:    tt.fields.Type,
				Method:  tt.fields.Method,
				Headers: tt.fields.Headers,
				Error:   tt.fields.Error,
			}
			if got := c.AuthorizationBearer(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Demand.AuthorizationBearer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemand_Authorization(t *testing.T) {
	type fields struct {
		URI     net_url.URL
		Token   string
		Type    string
		Method  string
		Headers map[string]string
		Error   error
	}
	type args struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Demand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Demand{
				URI:     tt.fields.URI,
				Token:   tt.fields.Token,
				Type:    tt.fields.Type,
				Method:  tt.fields.Method,
				Headers: tt.fields.Headers,
				Error:   tt.fields.Error,
			}
			if got := c.Authorization(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Demand.Authorization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemand_Header(t *testing.T) {
	type fields struct {
		URI     net_url.URL
		Token   string
		Type    string
		Method  string
		Headers map[string]string
		Error   error
	}
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Demand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Demand{
				URI:     tt.fields.URI,
				Token:   tt.fields.Token,
				Type:    tt.fields.Type,
				Method:  tt.fields.Method,
				Headers: tt.fields.Headers,
				Error:   tt.fields.Error,
			}
			if got := c.Header(tt.args.name, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Demand.Header() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemand_Parameter(t *testing.T) {
	type fields struct {
		URI     net_url.URL
		Token   string
		Type    string
		Method  string
		Headers map[string]string
		Error   error
	}
	type args struct {
		params any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Demand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Demand{
				URI:     tt.fields.URI,
				Token:   tt.fields.Token,
				Type:    tt.fields.Type,
				Method:  tt.fields.Method,
				Headers: tt.fields.Headers,
				Error:   tt.fields.Error,
			}
			if got := c.Parameter(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Demand.Parameter() = %v, want %v", got, tt.want)
			}
		})
	}
}
