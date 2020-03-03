package practice

import (
	"fmt"
	"testing"
)

// using diff name due to conflict
func TestConcatSubstring(t *testing.T) {

	s := "barfoofoobarthefoobarman"
	// s := "barfoothefoobarman"

	w := []string{"foo", "bar", "the"}

	fmt.Println(findSubstring(s, w))

}
