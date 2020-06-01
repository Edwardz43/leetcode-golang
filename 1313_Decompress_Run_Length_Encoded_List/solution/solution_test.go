package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/1313_Decompress_Run_Length_Encoded_List/solution"
	"github.com/stretchr/testify/assert"
)

func TestDecompressRLElist01(t *testing.T) {
	input := []int{1, 2, 3, 4}
	output := solution.DecompressRLElist(input)

	expexted := []int{2, 4, 4, 4}

	assert.Equal(t, output, expexted)
}

func TestDecompressRLElist02(t *testing.T) {
	input := []int{1, 1, 2, 3}
	output := solution.DecompressRLElist(input)

	expexted := []int{1, 3, 3}

	assert.Equal(t, output, expexted)
}
