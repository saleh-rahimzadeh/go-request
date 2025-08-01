package request

import (
	"io"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		timeout time.Duration
		retries []time.Duration
	}
	tests := []struct {
		name string
		args args
		want Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.timeout, tt.args.retries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_request_SendJson(t *testing.T) {
	type fields struct {
		Timeout time.Duration
		Retries []time.Duration
	}
	type args struct {
		c    Demand
		data any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Result
		want1  Properties
		want2  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := request{
				Timeout: tt.fields.Timeout,
				Retries: tt.fields.Retries,
			}
			got, got1, got2 := r.SendJson(tt.args.c, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("request.SendJson() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("request.SendJson() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("request.SendJson() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_request_SendForm(t *testing.T) {
	type fields struct {
		Timeout time.Duration
		Retries []time.Duration
	}
	type args struct {
		c    Demand
		data map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Result
		want1  Properties
		want2  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := request{
				Timeout: tt.fields.Timeout,
				Retries: tt.fields.Retries,
			}
			got, got1, got2 := r.SendForm(tt.args.c, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("request.SendForm() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("request.SendForm() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("request.SendForm() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_request_Send(t *testing.T) {
	type fields struct {
		Timeout time.Duration
		Retries []time.Duration
	}
	type args struct {
		c Demand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Result
		want1  Properties
		want2  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := request{
				Timeout: tt.fields.Timeout,
				Retries: tt.fields.Retries,
			}
			got, got1, got2 := r.Send(tt.args.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("request.Send() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("request.Send() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("request.Send() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_request_perform(t *testing.T) {
	type fields struct {
		Timeout time.Duration
		Retries []time.Duration
	}
	type args struct {
		c    Demand
		body io.Reader
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantResult    Result
		wantResponse  Properties
		wantIsSuccess bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := request{
				Timeout: tt.fields.Timeout,
				Retries: tt.fields.Retries,
			}
			gotResult, gotResponse, gotIsSuccess := r.perform(tt.args.c, tt.args.body)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("request.perform() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("request.perform() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotIsSuccess != tt.wantIsSuccess {
				t.Errorf("request.perform() gotIsSuccess = %v, want %v", gotIsSuccess, tt.wantIsSuccess)
			}
		})
	}
}

func Test_request_send(t *testing.T) {
	type fields struct {
		Timeout time.Duration
		Retries []time.Duration
	}
	type args struct {
		c    Demand
		body io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := request{
				Timeout: tt.fields.Timeout,
				Retries: tt.fields.Retries,
			}
			got, err := r.do(tt.args.c, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("request.send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("request.send() = %v, want %v", got, tt.want)
			}
		})
	}
}
