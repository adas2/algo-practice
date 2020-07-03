package strs

import (
	"fmt"
	"testing"

	custom "adas2.io/practice/custom_impl"
)

func setupTrie() *custom.Trie {
	words := []string{"to", "ton", "tons", "tip", "pit", "on"}
	dict := custom.InitTrie()

	for _, w := range words {
		dict.Insert(w)
	}

	return dict
}

func TestFindValidWord(t *testing.T) {
	g := [][]string{
		{"t", "i", "p"},
		{"o", "y", "g"},
		{"n", "s", "k"},
	}
	dict := setupTrie()
	findValidWords(g, dict)

	fmt.Println("match: ", dict.Find("tons"))
}
