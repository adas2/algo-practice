package trees

import (
	"fmt"
	"testing"
)

func TestInsertNode(t *testing.T) {
	bt := newBTree()
	bt.InsertNode(35)
	bt.InsertNode(5)
	bt.InsertNode(10)
	bt.InsertNode(73)
	bt.InsertNode(2)
	InorderPrintTree(bt.root)

}

func TestConstructBTree(t *testing.T) {
	inOrder := []int{1, 3, 5, 7, 10, 12}
	preOrder := []int{5, 3, 1, 10, 7, 12}
	tree, err := constructBTree(preOrder, inOrder)
	if err != nil {
		t.Fatalf("Contrusction failed %v", err)
	}
	fmt.Println("After Construction")
	InorderPrintTree(tree.root)
	tree.DeleteNode(5)
	InorderPrintTree(tree.root)
}

// func TestDeleteNode(t *testing.T){

// }
