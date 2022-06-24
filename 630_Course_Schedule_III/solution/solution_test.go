package solution

import "testing"

func Test_scheduleCourse(t *testing.T) {
	type args struct {
		courses [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "it should be 3",
			args: args{
				[][]int{
					{100, 200},
					{200, 1300},
					{1000, 1250},
					{2000, 3200},
				},
			},
			want: 3,
		},
		{
			name: "it should be 1",
			args: args{
				[][]int{
					{1, 2},
				},
			},
			want: 1,
		},
		{
			name: "it should be 0",
			args: args{
				[][]int{
					{3, 2},
					{4, 3},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scheduleCourse(tt.args.courses); got != tt.want {
				t.Errorf("scheduleCourse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkScheduleCourse(b *testing.B) {
	in := [][]int{
		{100, 200},
		{200, 1300},
		{1000, 1250},
		{2000, 3200},
	}
	for i := 0; i < b.N; i++ {
		scheduleCourse(in)
	}
}
func BenchmarkScheduleCourseV2(b *testing.B) {
	in := [][]int{
		{100, 200},
		{200, 1300},
		{1000, 1250},
		{2000, 3200},
	}
	for i := 0; i < b.N; i++ {
		scheduleCourseV2(in)
	}
}
