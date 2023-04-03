package trees

import (
	"fmt"
	"math"

	"adas2.io/practice/dp"
)

// Using order statitstics tree to keep node count
// TODO: add self balancing feature

// type osTree struct {
// 	root *osTreeNode
// }

const balanceTree = true

type osTreeNode struct {
	left   *osTreeNode
	right  *osTreeNode
	count  int
	value  int
	height int
}

// include count when inserting
func (root *osTreeNode) insert(val int) *osTreeNode {
	if root == nil {
		return &osTreeNode{
			left:   nil,
			right:  nil,
			value:  val,
			count:  1,
			height: 0,
		}
	}
	// not nil
	// BST property
	if root.value < val {
		root.right = root.right.insert(val)
	} else {
		root.left = root.left.insert(val)
	}
	// update height
	root.height = 1 + dp.Max(root.left.getHeight(), root.right.getHeight())
	// update count of root
	root.count++

	// if balancing is needed
	if balanceTree {
		// case: left-left
		if root.balanceFactor() > 1 && root.left.value > val {
			// root.rightRotate()
			return root.rightRotate()
		} else if root.balanceFactor() > 1 && root.left.value < val {
			// root.leftRotate() + root.rightRotate()
			root.left = root.left.leftRotate()
			return root.rightRotate()
		} else if root.balanceFactor() < -1 && root.right.value < val {
			// root.leftRotate()
			return root.leftRotate()
		} else if root.balanceFactor() < -1 && root.right.value > val {
			// root.rightRotate() + root.leftRotate()
			root.right = root.right.rightRotate()
			return root.leftRotate()
		}

	}

	return root
}

//print pre-order traversal of the binary search
func (root *osTreeNode) printTree() {
	if root != nil {
		// print root
		fmt.Printf("value: %d\tcount: %d\t height: %d\n", root.value, root.count, root.height)
		// print left subtree
		root.left.printTree()
		// print right subtree
		root.right.printTree()
	}
	// fmt.Println()
}

// find the kth smallest value (0-indexed)
// i.e 0th is the smallest, (n-1)th for largest
// return the node value where k = (root.count - root.right.count)
func (root *osTreeNode) findKthValue(k int) int {
	// kth value is is when: following condtion holds
	// count(root) - count(root.left) == k, return root.value
	// else recurse for subtress under root until above is true

	// error case: k cannot be negative
	if root == nil || k < 0 || k > root.count {
		return math.MaxInt32
	}

	lCount := root.left.getCount()

	// terminating condition:
	// when left subtree has k elements root is the k-th node
	if lCount == k {
		return root.value
	} else if k < lCount && root.left != nil {
		return root.left.findKthValue(k)
	}
	// else k > root.left.count
	if root.right != nil {
		return root.right.findKthValue(k - lCount - 1)
	}
	// all above cases don't match
	fmt.Printf("error: should not reach here:\n")
	return math.MaxInt32
}

// balance the os BST tree
func (root *osTreeNode) balanceFactor() int {
	if root == nil {
		return 0
	}
	// diff in the height of the left and right subtree
	return (root.left.getHeight() - root.right.getHeight())
}

// util
func (root *osTreeNode) getHeight() int {
	if root == nil {
		return 0
	}
	return root.height
}

func (root *osTreeNode) getCount() int {
	if root == nil {
		return 0
	}
	return root.count
}

func (root *osTreeNode) leftRotate() *osTreeNode {
	y := root.right
	t2 := y.left
	y.left = root
	root.right = t2

	// adjust height
	root.height = 1 + dp.Max(root.left.getHeight(), root.right.getHeight())
	y.height = 1 + dp.Max(y.left.getHeight(), y.right.getHeight())

	// adjust count?
	root.count = root.getCount() - y.getCount() + t2.getCount()
	y.count = y.getCount() - t2.getCount() + root.getCount()

	return y
}

func (root *osTreeNode) rightRotate() *osTreeNode {
	y := root.left
	t3 := y.right
	y.right = root
	root.left = t3

	// update height
	root.height = 1 + dp.Max(root.left.getHeight(), root.right.getHeight())
	y.height = 1 + dp.Max(y.left.getHeight(), y.right.getHeight())

	// adjust count? needed for OSTree
	root.count = root.getCount() - y.getCount() + t3.getCount()
	y.count = y.getCount() - t3.getCount() + root.getCount()

	return y
}
