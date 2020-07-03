package trees

import (
	"fmt"
	"sort"
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
	return &bTree{}
}

// InorderPrintTree :inorder traversal of the binary search
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
		return &btreeNode{val: aVal}
	}
	// evaluate based on a given property
	if evaluate(nPtr, aVal) == true {
		nPtr.left = nPtr.left.insert(aVal)
	} else {
		nPtr.right = nPtr.right.insert(aVal)
	}

	return nPtr
}

// DeleteNode searches nodes for a value and deletes it
// BST searching: assumes no duplicate values
func (t *bTree) DeleteNode(aVal int) {
	t.root = t.root.delete(aVal)
	return
}

// assumes unique values
func (nPtr *btreeNode) delete(aVal int) *btreeNode {
	// empty tree
	if nPtr == nil {
		return nil
	}

	// if value matches node val
	if aVal == nPtr.val {
		// case 1: leaf node: delete the node
		if nPtr.left == nil && nPtr.right == nil {
			nPtr = nil
		} else if nPtr.left == nil {
			// case 2a: right child only, or no child, swap child with node, delete node
			// nPtr.clearNode()
			nPtr = nPtr.right
		} else if nPtr.right == nil {
			// case 2b: left child only, or no child, swap child with node, delete node
			// nPtr.clearNode()
			nPtr = nPtr.left
		} else {
			// case 3: both children present, find next largest, i.e. leftmost in right subtree
			aPtr := nPtr.right.findLeftmost()
			nPtr.val = aPtr.val
			// tricky: successor may not a leaf, hence recursively delete the successor
			nPtr.right = nPtr.right.delete(aPtr.val)
		}

	} else if nPtr.val > aVal {
		nPtr.left = nPtr.left.delete(aVal)
	} else {
		nPtr.right = nPtr.right.delete(aVal)
	}

	return nPtr
}

// for memory safety: optional
func (nPtr *btreeNode) clearNode() {
	nPtr.val = -1   // safety
	nPtr.left = nil //safety
	nPtr.right = nil
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
func constructBTree(pot, iot []int) (*bTree, error) {
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

	if index = sort.SearchInts(iot, target); index >= len(iot) || iot[index] != target {
		return nil, fmt.Errorf("index %d node %d not found in IOT", index, target)
	}
	// value found: create new node corresponding to same
	node := &btreeNode{val: target, left: nil, right: nil}
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
// for BST (val <= root) go left else go right
func evaluate(node *btreeNode, val int) bool {
	return val <= node.val
}
