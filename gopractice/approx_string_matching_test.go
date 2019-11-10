package practice

import (
	"fmt"
	"testing"
)

func TestFindAllSubstringsMain(t *testing.T) {
	s := "tones"
	fmt.Println(FindSubtringPermutations(s))

}

func TestFindApproximateStringMatch(t *testing.T) {

	s := "stones"

	d := []string{"to", "toe", "note", "tone", "sons", "ones", "toner"}

	fmt.Println(FindApproximateStringMatch(d, s))
}
