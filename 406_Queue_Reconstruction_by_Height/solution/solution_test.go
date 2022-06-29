package solution

import (
	"reflect"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {
	type args struct {
		people [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				[][]int{
					{7, 0},
					{4, 4},
					{7, 1},
					{5, 0},
					{6, 1},
					{5, 2},
				},
			},
			want: [][]int{
				{5, 0},
				{7, 0},
				{5, 2},
				{6, 1},
				{4, 4},
				{7, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructQueue(tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

var in = [][]int{
	{7, 0},
	{4, 4},
	{7, 1},
	{5, 0},
	{6, 1},
	{5, 2},
}

func BenchmarkReconstructQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reconstructQueue(in)
	}
}
func BenchmarkReconstructQueueV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reconstructQueueV2(in)
	}
}
func BenchmarkReconstructQueueV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reconstructQueueV3(in)
	}
}
