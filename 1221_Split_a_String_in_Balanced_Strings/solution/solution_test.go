package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/1221_Split_a_String_in_Balanced_Strings/solution"
	"github.com/bmizerany/assert"
)

/*
	Example 1:
	Input: s = "RLRRLLRLRL"
	Output: 4

	Example 2:
	Input: s = "RLLLLRRRLR"
	Output: 3

	Example 3:
	Input: s = "LLLLRRRR"
	Output: 1

	Example 4:
	Input: s = "RLRRRLLRLL"
	Output: 2
*/
func TestBalancedStringSplit01(t *testing.T) {
	//TODO
	input := "RLRRLLRLRL"
	expexted := 4

	output := solution.BalancedStringSplit(input)

	assert.Equal(t, expexted, output)
}

func TestBalancedStringSplit02(t *testing.T) {
	//TODO
	input := "RLLLLRRRLR"
	expexted := 3

	output := solution.BalancedStringSplit(input)

	assert.Equal(t, expexted, output)
}

func TestBalancedStringSplit03(t *testing.T) {
	//TODO
	input := "LLLLRRRR"
	expexted := 1

	output := solution.BalancedStringSplit(input)

	assert.Equal(t, expexted, output)
}

func TestBalancedStringSplit04(t *testing.T) {
	//TODO
	input := "RLRRRLLRLL"
	expexted := 2

	output := solution.BalancedStringSplit(input)

	assert.Equal(t, expexted, output)
}
