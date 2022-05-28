package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/1476_Subrectangle_Queries/solution"
	"github.com/bmizerany/assert"
	"testing"
)

func TestSolution01(t *testing.T) {
	r := [][]int{
		[]int{1, 2, 1},
		[]int{4, 3, 4},
		[]int{3, 2, 1},
		[]int{1, 1, 1},
	}
	sq := solution.Constructor(r)
	expected := 1
	actual := sq.GetValue(0, 2)
	assert.Equal(t, expected, actual)
	sq.UpdateSubrectangle(0, 0, 3, 2, 5)
	expected = 5
	actual = sq.GetValue(0, 2)
	assert.Equal(t, expected, actual)
	expected = 5
	actual = sq.GetValue(3, 1)
	assert.Equal(t, expected, actual)
	sq.UpdateSubrectangle(3, 0, 3, 2, 10)
	expected = 5
	actual = sq.GetValue(0, 2)
	assert.Equal(t, expected, actual)
	expected = 10
	actual = sq.GetValue(3, 1)
	assert.Equal(t, expected, actual)
}
func TestSolution02(t *testing.T) {
	r := [][]int{
		[]int{1, 1, 1},
		[]int{2, 2, 2},
		[]int{3, 3, 3},
	}
	sq := solution.Constructor(r)
	expected := 1
	actual := sq.GetValue(0, 0)
	assert.Equal(t, expected, actual)
	sq.UpdateSubrectangle(0, 0, 2, 2, 100)
	expected = 100
	actual = sq.GetValue(0, 0)
	assert.Equal(t, expected, actual)
	expected = 100
	actual = sq.GetValue(2, 2)
	assert.Equal(t, expected, actual)
	sq.UpdateSubrectangle(1, 1, 2, 2, 20)
	expected = 20
	actual = sq.GetValue(2, 2)
	assert.Equal(t, expected, actual)
}
