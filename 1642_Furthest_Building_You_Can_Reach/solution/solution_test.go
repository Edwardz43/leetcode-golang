package solution

import "testing"

func Test_furthestBuilding(t *testing.T) {
	type args struct {
		heights []int
		bricks  int
		ladders int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "it should be 4",
			args: args{
				heights: []int{4, 2, 7, 6, 9, 14, 12},
				bricks:  5,
				ladders: 1,
			},
			want: 4,
		},
		{
			name: "it should be 7",
			args: args{
				heights: []int{4, 12, 2, 7, 3, 18, 20, 3, 19},
				bricks:  10,
				ladders: 2,
			},
			want: 7,
		},
		{
			name: "it should be 3",
			args: args{
				heights: []int{14, 3, 19, 3},
				bricks:  17,
				ladders: 0,
			},
			want: 3,
		},
		{
			name: "it should be 5",
			args: args{
				heights: []int{1, 5, 1, 2, 3, 4, 10000},
				bricks:  4,
				ladders: 1,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := furthestBuilding(tt.args.heights, tt.args.bricks, tt.args.ladders); got != tt.want {
				t.Errorf("furthestBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFurthestBuilding(b *testing.B) {
	heights := []int{1, 5, 1, 2, 3, 4, 10000}
	bricks := 4
	ladders := 1
	for i := 0; i < b.N; i++ {
		furthestBuilding(heights, bricks, ladders)
	}
}
func BenchmarkFurthestBuilding_v2(b *testing.B) {
	heights := []int{1, 5, 1, 2, 3, 4, 10000}
	bricks := 4
	ladders := 1
	for i := 0; i < b.N; i++ {
		furthestBuilding_v2(heights, bricks, ladders)
	}
}
