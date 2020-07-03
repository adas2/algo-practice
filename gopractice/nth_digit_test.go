package practice

import (
	"fmt"
	"testing"
)

func TestFindNthDigit(t *testing.T) {
	n := 19
	out := findNthDigit(n)
	fmt.Printf("input: %d digit: %d\n", n, out)
}
