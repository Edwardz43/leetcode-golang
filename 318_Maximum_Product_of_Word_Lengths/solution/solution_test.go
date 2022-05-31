package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/318_Maximum_Product_of_Word_Lengths/solution"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution01(t *testing.T) {
	input := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	expected := 16
	output := solution.MaxProduct(input)
	assert.Equal(t, expected, output)
}

var input = []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}

func BenchmarkMaxProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution.MaxProduct(input)
	}
}
func BenchmarkMaxProduct2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution.MaxProduct2(input)
	}
}
