package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/17_Letter_Combinations_of_a_Phone_Number/solution"
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "23",
			args: args{
				digits: "23",
			},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "empty",
			args: args{
				digits: "",
			},
			want: []string{},
		},
		{
			name: "2",
			args: args{
				digits: "2",
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "2345",
			args: args{
				digits: "2345",
			},
			want: []string{
				"adgj", "adgk", "adgl", "adhj", "adhk", "adhl", "adij", "adik", "adil", "aegj", "aegk", "aegl", "aehj",
				"aehk", "aehl", "aeij", "aeik", "aeil", "afgj", "afgk", "afgl", "afhj", "afhk", "afhl", "afij", "afik",
				"afil", "bdgj", "bdgk", "bdgl", "bdhj", "bdhk", "bdhl", "bdij", "bdik", "bdil", "begj", "begk", "begl",
				"behj", "behk", "behl", "beij", "beik", "beil", "bfgj", "bfgk", "bfgl", "bfhj", "bfhk", "bfhl", "bfij",
				"bfik", "bfil", "cdgj", "cdgk", "cdgl", "cdhj", "cdhk", "cdhl", "cdij", "cdik", "cdil", "cegj", "cegk",
				"cegl", "cehj", "cehk", "cehl", "ceij", "ceik", "ceil", "cfgj", "cfgk", "cfgl", "cfhj", "cfhk", "cfhl",
				"cfij", "cfik", "cfil",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.LetterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LetterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
