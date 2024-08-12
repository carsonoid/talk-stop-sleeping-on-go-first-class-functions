package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func Test_getURL(t *testing.T) {
	t.Skip("no way to reliably test this function")
}

type FakeTransport struct {
	response *http.Response
	err      error
}

func (t *FakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.response, t.err
}

func Test_getURLWithClient(t *testing.T) {
	type args struct {
		c   *http.Client
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				c: &http.Client{
					Transport: &FakeTransport{
						response: &http.Response{
							Body:       io.NopCloser(&bytes.Buffer{}),
							StatusCode: 200,
						},
					},
				},
			},
			want: http.StatusOK,
		},
		{
			name: "fail",
			args: args{
				c: &http.Client{
					Transport: &FakeTransport{
						err: io.EOF,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getURLWithClient(tt.args.c, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getURLWithClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getURLWithClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

type fakeGetter struct {
	response *http.Response
	err      error
}

func (f *fakeGetter) Get(url string) (resp *http.Response, err error) {
	return f.response, f.err
}

func Test_getURLWithInterface(t *testing.T) {
	type args struct {
		g   getter
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				g: &fakeGetter{
					response: &http.Response{
						Body:       io.NopCloser(&bytes.Buffer{}),
						StatusCode: 200,
					},
				},
			},
			want: http.StatusOK,
		},
		{
			name: "fail",
			args: args{
				g: &fakeGetter{
					err: io.EOF,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getURLWithInterface(tt.args.g, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getURLWithInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getURLWithInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getURLWithGlobal(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		args     args
		testFunc getFunc
		want     int
		wantErr  bool
	}{
		{
			name: "test",
			args: args{},
			testFunc: func(url string) (resp *http.Response, err error) {
				return &http.Response{
					Body:       io.NopCloser(&bytes.Buffer{}),
					StatusCode: 200,
				}, nil
			},
			want: http.StatusOK,
		},
		{
			name: "fail",
			args: args{},
			testFunc: func(url string) (resp *http.Response, err error) {
				return nil, io.EOF
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ensure cleanup of global variable
			cur := globalGetFunc
			defer func() {
				globalGetFunc = cur
			}()

			// override global variable
			globalGetFunc = tt.testFunc

			got, err := getURLWithGlobal(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getURLWithGlobal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getURLWithGlobal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getURLWithFunc(t *testing.T) {
	type args struct {
		get getFunc
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				get: func(url string) (resp *http.Response, err error) {
					return &http.Response{
						Body:       io.NopCloser(&bytes.Buffer{}),
						StatusCode: 200,
					}, nil
				},
			},
			want: 200,
		},
		{
			name: "fail",
			args: args{
				get: func(url string) (resp *http.Response, err error) {
					return nil, io.EOF
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getURLWithFunc(tt.args.get, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getURLWithFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getURLWithFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
