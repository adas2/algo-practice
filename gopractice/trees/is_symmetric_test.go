package trees

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	root := &btreeNode{
		1,
		&btreeNode{2, &btreeNode{3, nil, nil}, &btreeNode{4, &btreeNode{5, nil, nil}, nil}},
		&btreeNode{2, &btreeNode{4, nil, &btreeNode{5, nil, nil}}, &btreeNode{3, nil, nil}},
	}

	fmt.Println(isSymmetric(root, root))

	fmt.Printf("Iterative version: %v\n", isSymmetricIterative(root))

}
