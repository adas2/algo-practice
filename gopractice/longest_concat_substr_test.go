package practice

import (
	"fmt"
	"testing"
)

// using diff name due to conflict
func TestConcatSubstring(t *testing.T) {

	// s := "barfoofoobarthefoobarman"
	// w := []string{"foo", "bar", "the"}
	// fmt.Println(findSubstring(s, w))

	s2 := "wordgoodgoodgoodbestword"
	w2 := []string{"word", "good", "best", "good"}
	fmt.Println(findSubstring(s2, w2))

}
