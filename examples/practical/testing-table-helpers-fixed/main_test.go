package main

import (
	"context"
	"testing"
)

// START ORIGINAL OMIT
func Test_getHello(t *testing.T) {
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
			name: "cancelled",
			args: args{
				ctx: func() context.Context { // 🤩 create and call in the table!
					cancelledCtx, done := context.WithCancel(context.Background())
					done()
					return cancelledCtx
				}(),
			},
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
