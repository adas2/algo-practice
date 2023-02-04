package trees

// given a binary tree where each node has an int.
// path weight of a node is the sum of the int vals on the
// unique path from the root to the node.
// givem a value check if there is exists a node with given path weight

func checkValidPathWeight(root *btreeNode, sum int) bool {
	// null subtree means no path foun d
	if root == nil {
		return false
	}

	// match found
	if sum == root.val {
		return true
	}

	// else recurse left and right subtree
	return checkValidPathWeight(root.left, sum-root.val) ||
		checkValidPathWeight(root.right, sum-root.val)

}

// logic:
// use recursion at every level pass the diff = (sum-node.val)
// to the left and right subtrees. If the difference is zero for a node, return true
// else false
