package practice

import "testing"

func TestFindMaxValuePath(t *testing.T) {
	s := [][]int{
		{1, 3, 7},
		{4, 2, 5},
		{0, 6, 10},
	}

	FindMaxValuePath(s)
}
