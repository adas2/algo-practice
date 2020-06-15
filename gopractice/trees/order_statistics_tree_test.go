package trees

import (
	"fmt"
	"testing"
)

func TestOSTree(t *testing.T) {
	var tr *osTreeNode
	// tr := nil
	tr = tr.insert(10)
	tr = tr.insert(5)
	tr = tr.insert(15)
	tr = tr.insert(35)
	tr = tr.insert(20)

	tr.printTree()
	fmt.Println("balance-factor: ", tr.balanceFactor())
	k := 2
	fmt.Println("k:", k, "value:", tr.findKthValue(k))

}
