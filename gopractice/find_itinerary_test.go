package practice

import (
	"fmt"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	tix := [][]string{
		// {"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"},
		{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"},
	}
	res := findItinerary(tix)
	fmt.Println(res)
}
