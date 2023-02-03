package trees

import (
	"os"
	"testing"
)

func TestFindKBalancedMain(t *testing.T) {

	// syntehtic tree find a better a insert function
	bt := newBTree()
	bt.InsertNode(50)
	bt.InsertNode(20)
	bt.InsertNode(60)
	bt.InsertNode(10)
	bt.InsertNode(25)
	bt.InsertNode(55)
	bt.InsertNode(70)
	bt.InsertNode(65)
	bt.InsertNode(80)
	bt.InsertNode(85)
	bt.InsertNode(90)

	PrintLevelOrder(os.Stdout, bt.root)

	findKbalancedMain(bt.root)

}
