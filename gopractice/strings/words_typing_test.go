package strs

import (
	"fmt"
	"testing"
)

func TestWordTyping(t *testing.T) {
	s := []string{"i", "had", "apple", "pie"}
	fmt.Println("Input", s)
	fmt.Println(wordsTyping(s, 4, 5))
}
