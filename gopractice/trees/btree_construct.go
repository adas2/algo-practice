package trees

import (
	"fmt"

	"adas2.io/practice"
)

type btreeNode struct {
	val   int
	left  *btreeNode
	right *btreeNode
}

type bTree struct {
	root *btreeNode
}

// tree cobstructor
func newBTree() *bTree {
	return &bTree{root: nil}
}

// Methods:
// inorder traversal of the binary search
func InorderPrintTree(root *btreeNode) {
	if root != nil {
		// print left subtree
		InorderPrintTree(root.left)
		// print root
		fmt.Printf("%d ", root.val)
		// print right subtree
		InorderPrintTree(root.right)
	}
}

// InsertNode inserts a new node into bTree t
// wrapper for bNode insert()
func (t *bTree) InsertNode(aVal int) /**bTree*/ {
	t.root = t.root.insert(aVal)
	return
}

// insert a value at a given node/root, return update node/root
func (nPtr *btreeNode) insert(aVal int) *btreeNode {
	if nPtr == nil {
		// create node
		return &btreeNode{val: aVal, left: nil, right: nil}
	}
	// evaluate based on a given property
	if evaluate(nPtr, aVal) == true {
		nPtr.left = nPtr.left.insert(aVal)
	} else {
		nPtr.right = nPtr.right.insert(aVal)
	}

	return nPtr
}

// DeleteNode searches the value and deletes it
// BST searching: assumes no duplicate values
func (t *bTree) DeleteNode(aVal int) {
	t.root = t.root.delete(aVal)
	return
}

func (nPtr *btreeNode) delete(aVal int) *btreeNode {
	// empty tree
	if nPtr == nil {
		return nil
	}

	// if value matches node vale
	if aVal == nPtr.val {
		fmt.Println("Match value:", aVal)
		// case 1: leaf node: delete the node
		if nPtr.left == nil && nPtr.right == nil {
			nPtr = nil
		} else if nPtr.left == nil && nPtr.right != nil {
			// case 2a: one child only, swap child with node, delete node
			nPtr.clearNode()
			nPtr = nPtr.right
		} else if nPtr.right == nil && nPtr.left != nil {
			// case 2b: one child only, swap child with node, delete node
			nPtr.clearNode()
			nPtr = nPtr.left
		} else {
			// case 3: both children present, find next largest, i.e. leftmost in right subtree
			aPtr := nPtr.right.findLeftmost()
			// note successor in this case is always a leaf
			nPtr = aPtr
			// aPtr = nil
		}

	}

	// parent = nPtr
	if nPtr.val > aVal {
		nPtr = nPtr.left.delete(aVal)
	} else {
		nPtr = nPtr.right.delete(aVal)
	}
	// else { //equal

	// }
	return nPtr
}

func (nPtr *btreeNode) clearNode() {
	nPtr.val = -1   // safety
	nPtr.left = nil //safety
	nPtr.right = nil
	return
}

// helper to find the next largest element
func (nPtr *btreeNode) findLeftmost() *btreeNode {
	// leftmost child of the right subtree
	if nPtr.left != nil {
		nPtr.left.findLeftmost()
	}

	return nPtr
}

// ConstructBTree creates a binary tree from inorder (iot) and preorder (pot) traversal
func ConstructBTree(pot, iot []int) (*bTree, error) {
	// must have equal num non-zero elements
	if len(iot) != len(pot) {
		return nil, fmt.Errorf("Invalid input")
	}

	bt, err := constructbNode(pot, iot, 0, len(pot)-1)

	return &bTree{root: bt}, err
}

// helper function for ConstructBTree
func constructbNode(pot, iot []int, li, ri int) (*btreeNode, error) {
	if len(iot) == 0 || len(pot) == 0 {
		return nil, nil
	}

	fmt.Println(pot, iot, li, ri)
	// search pot[li] in iot
	var index int
	var err error
	target := pot[li]
	if index = practice.BinarySearchVanilla(iot, target, 0, len(iot)-1); index == -1 {
		return nil, fmt.Errorf("node %d not found in Inorder Traversal list", target)
	}
	// value found: create new node corresponding to same
	node := &btreeNode{val: pot[li], left: nil, right: nil}
	node.left, err = constructbNode(pot, iot[:index], li+1, li+index)
	if err != nil {
		return nil, err
	}
	node.right, err = constructbNode(pot, iot[(index+1):], li+index+1, ri)
	if err != nil {
		return nil, err
	}

	return node, nil

}

// evaluate property of tree
// distinguishes type of tree e.g. BST has property if (val <= root) go left else go right
func evaluate(node *btreeNode, val int) bool {
	if val <= node.val {
		return true
	}
	return false
}