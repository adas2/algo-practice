package practice

import (
	"fmt"
	"testing"
)

func TestMaxEnvelops(t *testing.T) {
	envs := [][]int{
		{6, 4}, {6, 7}, {2, 2}, {3, 3}, {3, 4}, {3, 5}, {4, 5},
	}
	fmt.Println("Max evelops: ", maxEnvelopes(envs))
}
