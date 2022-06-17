package solution

import (
	"reflect"
	"testing"
)

func Test_getConcatenation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				[]int{1, 2, 1},
			},
			want: []int{1, 2, 1, 1, 2, 1},
		},
		{
			name: "test 2",
			args: args{
				[]int{1, 3, 2, 1},
			},
			want: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConcatenation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConcatenation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetConcatenation(b *testing.B) {
	in := []int{1, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		getConcatenation(in)
	}
}
func BenchmarkGetConcatenation_v2(b *testing.B) {
	in := []int{1, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		getConcatenation_v2(in)
	}
}
