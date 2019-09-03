package practice

import (
	"os"
	"testing"
)

func TestPrintLevelOrder(t *testing.T) {
	bt := newBTree()
	bt.InsertNode(35)
	bt.InsertNode(5)
	bt.InsertNode(10)
	bt.InsertNode(73)
	bt.InsertNode(2)

	// InorderPrintTree(bt.root)
	PrintLevelOrder(os.Stdout, bt)
}
