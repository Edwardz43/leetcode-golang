package solution

import "testing"

func Test_maximumUnits(t *testing.T) {
	type args struct {
		boxTypes  [][]int
		truckSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "it should be 8",
			args: args{
				boxTypes: [][]int{
					{1, 3}, {2, 2}, {3, 1},
				},
				truckSize: 4,
			},
			want: 8,
		},
		{
			name: "it should be 91",
			args: args{
				boxTypes: [][]int{
					{5, 10}, {2, 5}, {4, 7}, {3, 9},
				},
				truckSize: 10,
			},
			want: 91,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUnits(tt.args.boxTypes, tt.args.truckSize); got != tt.want {
				t.Errorf("maximumUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}
