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

func TestFindSubstrings(t *testing.T) {
	s := "panama"

	fmt.Println(FindSubstrings(s))
}

func TestFindClosestWord(t *testing.T) {

	s := "stones"

	d := []string{"to", "toe", "note", "tone", "sons", "toner"}

	fmt.Println(FindClosestWord(s, d))

}
