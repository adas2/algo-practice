package strs

import (
	"fmt"
	"testing"
)

func TestFindPattern(t *testing.T) {
	text := "AABABABBABAC"
	pat := "ABA"

	// r := []rune(text)
	// fmt.Printf("rune: %v, int_val: %v, ", int(r[0]), int(text[0]))

	fmt.Println("text:", text, "pattern:", pat, "len_str:", len(text))
	res := findPattern(text, pat)
	fmt.Println("matches:", res)
}
