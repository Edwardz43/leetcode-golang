package solution

import "testing"

func Test_minDeletions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "it should 0",
			args: args{
				s: "aab",
			},
			want: 0,
		},
		{
			name: "it should 0",
			args: args{
				s: "aab",
			},
			want: 0,
		},
		{
			name: "it should 2",
			args: args{
				s: "ceabaacb",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletions(tt.args.s); got != tt.want {
				t.Errorf("minDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
