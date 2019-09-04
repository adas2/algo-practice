package practice

import "fmt"

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

// indert a value at a given node
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
func (*bTree) DeleteNode(aVal int) error {
	return nil
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
	if index = BinarySearchVanilla(iot, target, 0, len(iot)-1); index == -1 {
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
