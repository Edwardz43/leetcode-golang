package solution

import (
	"testing"
)

func Test_longestStrChain(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "it should be 4",
			args: args{
				[]string{
					"a", "b", "ba", "bca", "bda", "bdca",
				},
			},
			want: 4,
		},
		{
			name: "it should be 5",
			args: args{
				[]string{
					"xbc", "pcxbcf", "xb", "cxbc", "pcxbc",
				},
			},
			want: 5,
		},
		{
			name: "it should be 1",
			args: args{
				[]string{
					"abcd", "dbqca",
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestStrChain_v2(tt.args.words); got != tt.want {
				t.Errorf("longestStrChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_longestStrChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestStrChain([]string{
			"a", "b", "ba", "bca", "bda", "bdca",
		})
	}
}
func Benchmark_longestStrChain_v2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestStrChain_v2([]string{
			"a", "b", "ba", "bca", "bda", "bdca",
		})
	}
}
func Benchmark_longestStrChain_v3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestStrChain_v3([]string{
			"a", "b", "ba", "bca", "bda", "bdca",
		})
	}
}
