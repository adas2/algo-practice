package trees

import (
	"fmt"
	"testing"
)

func TestCheckIfBst(t *testing.T) {
	// sample input
	// Level Order = [2,2,2]
	// 			2
	//      2 		2

	bt := &btreeNode{2, &btreeNode{2, nil, nil}, &btreeNode{2, nil, nil}}

	InorderPrintTree(bt)

	fmt.Println("Is BST? ", checkIfBst2(bt))

}
