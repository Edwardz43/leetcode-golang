package solution

import (
	"testing"
)

func TestSolution01(t *testing.T) {
	// input :=
	// expected :=
	// output := solution.function(input)
	// assert.Equal(t, expected, output)
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "test2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "test3",
			args: args{
				s: "abvghjkjhgqwe",
			},
			want: "ghjkjhg",
		},
		{
			name: "test4",
			args: args{
				s: "baabad",
			},
			want: "baab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome_v2(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_longestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestPalindrome("abvghjkjhgqwe")
	}
}
func Benchmark_longestPalindrome_v2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestPalindrome_v2("abvghjkjhgqwe")
	}
}
