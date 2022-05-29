package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/14_Longest_Common_Prefix/solution"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution01(t *testing.T) {
	input := []string{"flower", "flow", "flight"}
	expected := "fl"
	output := solution.LongestCommonPrefix(input)
	assert.Equal(t, expected, output)
}

func TestSolution02(t *testing.T) {
	input := []string{"dog", "racecar", "car"}
	expected := ""
	output := solution.LongestCommonPrefix(input)
	assert.Equal(t, expected, output)
}

func TestSolution03(t *testing.T) {
	input := []string{"a"}
	expected := "a"
	output := solution.LongestCommonPrefix(input)
	assert.Equal(t, expected, output)
}

func TestSolution04(t *testing.T) {
	input := []string{"", ""}
	expected := ""
	output := solution.LongestCommonPrefix(input)
	assert.Equal(t, expected, output)
}

func TestSolution05(t *testing.T) {
	input := []string{"a", ""}
	expected := ""
	output := solution.LongestCommonPrefix(input)
	assert.Equal(t, expected, output)
}

func TestSolution06(t *testing.T) {
	input := []string{"flower", "flower", "flower", "flower"}
	expected := "flower"
	output := solution.LongestCommonPrefix(input)
	assert.Equal(t, expected, output)
}
