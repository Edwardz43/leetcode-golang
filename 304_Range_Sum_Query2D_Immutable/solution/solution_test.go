package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/304_Range_Sum_Query2D_Immutable/solution"
	"testing"
)

func TestNumMatrix_SumRegion(t *testing.T) {
	type fields struct {
		matrix [][]int
	}
	f := fields{
		matrix: [][]int{
			{3, 0, 1, 4, 2},
			{5, 6, 3, 2, 1},
			{1, 2, 0, 1, 5},
			{4, 1, 0, 1, 7},
			{1, 0, 3, 0, 5},
		},
	}
	type args struct {
		row1 int
		col1 int
		row2 int
		col2 int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "it should be 8 while given arr = [2, 1, 4, 3]",
			fields: f,
			args: args{
				row1: 2,
				col1: 1,
				row2: 4,
				col2: 3,
			},
			want: 8,
		},
		{
			name:   "it should be 11 while given arr = [1, 2, 2, 2]",
			fields: f,
			args: args{
				row1: 1,
				col1: 1,
				row2: 2,
				col2: 2,
			},
			want: 11,
		},
		{
			name:   "it should be 12 while given arr = [1, 2, 2, 4]",
			fields: f,
			args: args{
				row1: 1,
				col1: 2,
				row2: 2,
				col2: 4,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := solution.Constructor(tt.fields.matrix)
			if got := this.SumRegion(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2); got != tt.want {
				t.Errorf("SumRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}
