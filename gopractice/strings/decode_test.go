package strs

import (
	"fmt"
	"testing"
)

func TestDecodeString(t *testing.T) {
	s := "3[a2[c]]"

	fmt.Println(decodeString(s))
}
