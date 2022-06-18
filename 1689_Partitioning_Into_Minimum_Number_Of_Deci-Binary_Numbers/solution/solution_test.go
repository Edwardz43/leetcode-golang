package solution

import "testing"

func Test_minPartitions(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "it should be 3",
			args: args{
				n: "32",
			},
			want: 3,
		},
		{
			name: "it should be 8",
			args: args{
				n: "82734",
			},
			want: 8,
		},
		{
			name: "it should be 9",
			args: args{
				n: "27346209830709182346",
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPartitions(tt.args.n); got != tt.want {
				t.Errorf("minPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

const in = "27346209830709182346"

func BenchmarkMinPartitions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minPartitions(in)
	}
}
func BenchmarkMinPartitions_v2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minPartitions_v2(in)
	}
}
