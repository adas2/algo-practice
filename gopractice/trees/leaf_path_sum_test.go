package trees

import (
	"fmt"
	"os"
	"testing"
)

func TestFindLeafPathSum(t *testing.T) {
	bt := &btreeNode{1, &btreeNode{0, &btreeNode{0, nil, nil}, &btreeNode{1, nil, nil}}, &btreeNode{1, nil, nil}}

	PrintLevelOrder(os.Stdout, bt)

	fmt.Println("Leaf path sum :", findLeafPathSum(bt))
}

// func TestAddBitst(t *testing.T) {
// 	a := NewBitNum()
// 	b := NewBitNum()
// 	a = AppenBit(AppenBit(AppenBit(a, 1), 0), 1)
// 	b = AppenBit(b, 1)
// 	PrintBits(a)
// 	PrintBits(b)
// 	result, carry := AddBits(a, b, 0)
// 	PrintBits(AppenBit(result, carry))
// }
