package strs

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	s := "bcabc"

	out := removeDuplicateLetters(s)
	fmt.Println(out)
}
