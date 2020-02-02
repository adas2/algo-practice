package practice

import (
	"fmt"
	"testing"
)

func TestSortSimilarItems(t *testing.T) {
	Inp := []string{"L", "M", "S", "L", "L", "S", "M", "M"}
	fmt.Println(Inp)
	SortSimilarItems(Inp)
}
