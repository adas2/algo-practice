package strs

import (
	"fmt"
	"testing"
)

func TestGroupStrings(t *testing.T) {
	strs := []string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}
	// strs := []string{"a"}
	fmt.Println("Input", strs)
	fmt.Println(groupStrings(strs))
}
