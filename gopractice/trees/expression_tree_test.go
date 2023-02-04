package trees

import (
	"fmt"
	"testing"
)

func TestCheckEquivalence(t *testing.T) {
	b1 := &extreeNode{"+", &extreeNode{"a", nil, nil},
		&extreeNode{"+", &extreeNode{"b", nil, nil}, &extreeNode{"c", nil, nil}}}

	b2 := &extreeNode{"+", &extreeNode{"+", &extreeNode{"b", nil, nil}, &extreeNode{"d", nil, nil}},
		&extreeNode{"a", nil, nil}}

	fmt.Println(checkEquivalence(b1, b2))

}
