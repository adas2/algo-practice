package strs

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}

	fmt.Println(wordBreakDp(s, wordDict))
}
