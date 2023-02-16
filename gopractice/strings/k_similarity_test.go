package strs

import (
	"fmt"
	"testing"
)

func TestKSimilarity(t *testing.T) {
	s1, s2 := "abccaacceecdeea", "bcaacceeccdeaae"
	fmt.Printf("Input: %s, %s\n", s1, s2)
	fmt.Println(kSimilarity(s1, s2))
}
