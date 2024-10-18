package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_getNextHours(t *testing.T) {
	type args struct {
		n int
	}
	// START TESTS OMIT
	tests := []struct {
		name   string
		args   args
		want   []time.Time
		assert func(*testing.T, []time.Time) // 🎉
	}{
		// END TESTS OMIT
		// START CASE OMIT
		{
			name: "0 days",
			args: args{n: 0},
			want: nil,
		},
		{
			name: "3 days",
			args: args{n: 3},
			assert: func(t *testing.T, s []time.Time) {
				if len(s) != 3 {
					t.Errorf("expected 3 days, got %d", len(s))
				}

				for i, d := range s {
					if d.Before(time.Now()) {
						t.Errorf("expected future date, got %v", d)
					}

					if i > 0 && d.Sub(s[i-1]) != time.Hour {
						t.Errorf("expected 1 hour difference, got %v", d.Sub(s[i-1]))
					}
				}
			},
		},
		// END CASE OMIT
	}

	// START RUN OMIT
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getNextHours(tt.args.n)

			if tt.assert != nil {
				tt.assert(t, got)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNextHours() = %v, want %v", got, tt.want)
			}
		})
	}
	// END RUN OMIT
}
