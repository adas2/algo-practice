package practice

import (
	"fmt"
	"testing"
)

func TestFindMissingLetters(t *testing.T) {
	s := "The quick brown fox jumps over the Lay Do"

	l := FindMissingLetters(s)

	fmt.Println(l)
}

func TestFindDictWord(t *testing.T) {
	s := "A man a plan a canal panama"

	d := []string{"man", "plan", "canal"}

	fmt.Println(FindDictWord(s, d))
}
