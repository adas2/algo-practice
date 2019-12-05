package practice

import (
	"fmt"
	"testing"
)

func TestFindRotationPoint(t *testing.T) {

	// sList := []string{}
	sList := []string{"dinosaur", "elephant", "horse", "zebra", "apple", "banana", "coconut"}
	// sList := []string{"apple", "banana", "cat", "coconut", "dinosaur", "elephant", "horse", "zebra"}

	fmt.Printf("The rotation point for the array %q is %d\n", sList, findRotationPoint(sList))

}
