package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_getNextHoursFn(t *testing.T) {
	// START TESTS OMIT
	type args struct {
		getNow func() time.Time
		n      int
	}
	tests := []struct {
		name string
		args args
		want []time.Time
	}{
		// END TESTS OMIT
		// START CASE OMIT
		{
			name: "0 hours",
			args: args{
				n:      0,
				getNow: func() time.Time { return time.Time{} },
			},
			want: nil,
		},
		{
			name: "3 hours",
			args: args{
				n:      3,
				getNow: func() time.Time { return time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC) },
			},
			want: []time.Time{
				time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 1, 2, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 1, 3, 0, 0, 0, time.UTC),
			},
		},
		// END CASE OMIT
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNextHoursFn(tt.args.getNow, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNextHoursFn() = %v, want %v", got, tt.want)
			}
		})
	}
}

// END ORIGINAL OMIT
