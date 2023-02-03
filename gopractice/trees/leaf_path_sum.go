package trees

// given a binary tree where each node has a val of 0/1,
// leaf path value is the integer value for the binary number with MSB as root val
// and LSB as the leaf val. Find the sum of all such leaf path integer values
// e.g.					1
//					0 /    \ 1
// 				0/ 	  \1
// leaf path sum = (100)2 + (101)2 + (11)2 = 4 + 5 + 3 = 12

// Trick: Use recursive util to calculate sum of int values up to that node
func findLeafPathSum(root *btreeNode) int {

	// int value for parent is initially zero
	return leafPathSumHelper(root, 0)
}

// helper function: uses the int value representation at parent node
// it represents the partial sum of the bits faced till the parent node
func leafPathSumHelper(root *btreeNode, nodeIntValue int) int {

	// find the integer value at the given node with the help of it parent node

	// if node is null, node_int_value = 0
	if root == nil {
		return 0
	}

	// (parent_int_value * 2) + node_value = node_int_value
	nodeIntValue = 2*nodeIntValue + root.val

	// if node is leaf, return node_int_value, which is the leaf sum path
	if root.left == nil && root.right == nil {
		return nodeIntValue
	}

	// non-leaf
	// the sum is node_int_value(left_node) + node_int_value(right_node)
	return leafPathSumHelper(root.right, nodeIntValue) +
		leafPathSumHelper(root.left, nodeIntValue)

}
