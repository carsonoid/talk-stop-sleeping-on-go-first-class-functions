package main

import (
	"context"
	"testing"
)

// START ORIGINAL OMIT
func Test_getHello(t *testing.T) {
	// 🤮 - code outside the test!
	cancelledCtx, done := context.WithCancel(context.Background())
	done()

	type args struct{ ctx context.Context }
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "not cancelled",
			args: args{ctx: context.Background()},
			want: "Hello, world!",
		},
		{
			name:    "cancelled",
			args:    args{ctx: cancelledCtx /*🤮*/},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getHello(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("getHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getHello() = %v, want %v", got, tt.want)
			}
		})
	}
}

// END ORIGINAL OMIT
