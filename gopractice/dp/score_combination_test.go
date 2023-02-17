package dp

import (
	"fmt"
	"testing"
)

func TestScoreCobination(t *testing.T) {
	f := 12
	p := []int{2, 3, 7, 8}

	fmt.Println(scoreCombinations(f, p))
}
