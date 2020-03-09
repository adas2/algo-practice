package practice

import (
	"fmt"
	"testing"
)

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	dict := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}

	r := findAllConcatenatedWordsInADict(dict)

	fmt.Println(r)
}
