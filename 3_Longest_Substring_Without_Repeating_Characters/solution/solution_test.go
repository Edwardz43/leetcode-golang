package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/3_Longest_Substring_Without_Repeating_Characters/solution"

	"github.com/bmizerany/assert"
)

func TestSolution01(t *testing.T) {
	input := "abcabcbb"
	expected := 3
	output := solution.LengthOfLongestSubstring(input)
	assert.Equal(t, expected, output)
}

func TestSolution02(t *testing.T) {
	input := "bbbbb"
	expected := 1
	output := solution.LengthOfLongestSubstring(input)
	assert.Equal(t, expected, output)
}

func TestSolution03(t *testing.T) {
	input := "pwwkew"
	expected := 3
	output := solution.LengthOfLongestSubstring(input)
	assert.Equal(t, expected, output)
}

func TestSolution04(t *testing.T) {
	input := ""
	expected := 0
	output := solution.LengthOfLongestSubstring(input)
	assert.Equal(t, expected, output)
}

func TestSolution05(t *testing.T) {
	input := "abcabc"
	expected := 3
	output := solution.LengthOfLongestSubstring(input)
	assert.Equal(t, expected, output)
}

func TestSolution06(t *testing.T) {
	input := "dvdf"
	expected := 3
	output := solution.LengthOfLongestSubstring(input)
	assert.Equal(t, expected, output)
}

func TestSolution07(t *testing.T) {
	input := "au"
	expected := 2
	output := solution.LengthOfLongestSubstring(input)
	assert.Equal(t, expected, output)
}
