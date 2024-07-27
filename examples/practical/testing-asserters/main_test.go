package main

import (
	"reflect"
	"testing"
	"time"
)

// START ORIGINAL OMIT
func Test_getNextHours(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []time.Time
	}{
		{
			name: "0 hours",
			args: args{n: 0},
			want: nil,
		},
		{
			name: "3 hours",
			args: args{n: 3},
			want: []time.Time{
				// ???
			},
		},
	}
	// END ORIGINAL OMIT
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNextHours(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNextHours() = %v, want %v", got, tt.want)
			}
		})
	}
}
