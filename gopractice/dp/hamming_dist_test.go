package dp

import (
	"fmt"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	s1 := "geeksforgeeks"
	s2 := "geeksandgeeks"
	s3 := "geeksorgeeks"

	fmt.Println(HammingDistance(s1, s2))
	fmt.Println(HammingDistance(s1, s3))

}

func TestEditDistance(t *testing.T) {
	s1 := "anthem"
	s2 := "anthurium"
	s3 := "antmen"

	fmt.Println(EditDistance(s1, s2))
	fmt.Println(EditDistance(s1, s3))
}
