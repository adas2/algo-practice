package practice

import (
	"fmt"
	"testing"
)

func TestCommandCount(t *testing.T) {

	inp := []string{"abc:/b1c\\xy", "w::/a\\b"}

	fmt.Println(commandCount(inp))

	// fmt.Println(getSubstrings(inp[1]))

}
