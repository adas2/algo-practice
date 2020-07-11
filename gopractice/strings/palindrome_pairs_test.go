package strs

import (
	"fmt"
	"testing"
)

func TestCheckPalindrome(t *testing.T) {
	str := "b"

	fmt.Println(checkPalindrome(str, 0, len(str)-1))

	fmt.Println(reverseWord(str))
}

func TestPalindromePairs(t *testing.T) {
	words := []string{"abcd", "dcba", "lls", "s", "sssll"}

	res := palindromePairs(words)

	fmt.Println(res)
}
