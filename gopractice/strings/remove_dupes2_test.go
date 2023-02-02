package strs

import (
	"fmt"
	"testing"
)

func TestRemoveDupes(t *testing.T) {
	str := "pbbcggttciiippooaais"
	k := 2

	fmt.Println("After removal:", removeDuplicates(str, k))
}
