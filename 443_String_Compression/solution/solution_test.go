package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/443_String_Compression/solution"
	"github.com/bmizerany/assert"
)

func TestCompress01(t *testing.T) {
	input := []byte("aaabbccc")
	output := solution.Compress(input)

	assert.Equal(t, 6, output)
}

func TestCompress02(t *testing.T) {
	input := []byte("aabcccccccccc")
	output := solution.Compress(input)

	assert.Equal(t, 6, output)
}

func TestCompress03(t *testing.T) {
	input := []byte("a")
	output := solution.Compress(input)

	assert.Equal(t, 1, output)
}
