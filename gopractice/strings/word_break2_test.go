package strs

import (
	"fmt"
	"testing"
)

func TestWordBreak2(t *testing.T) {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}

	out := wordBreak2(s, wordDict)
	for _, s := range out {
		fmt.Printf("%s<>\n", s)
	}
}
