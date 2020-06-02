package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/1295_Find_Numbers_with_Even_Number_of_Digits/solution"
	"github.com/bmizerany/assert"
)

func TestFindNumbers01(t *testing.T) {
	input := []int{12, 345, 2, 6, 7896}
	expexted := 2
	output := solution.FindNumbers(input)

	assert.Equal(t, expexted, output)
}

func TestFindNumbers02(t *testing.T) {
	input := []int{555, 901, 482, 1771}
	expexted := 1
	output := solution.FindNumbers(input)

	assert.Equal(t, expexted, output)
}

func TestFindNumbers03(t *testing.T) {
	input := []int{100000, 30000}
	expexted := 0
	output := solution.FindNumbers(input)

	assert.Equal(t, expexted, output)
}
