package trees

import "math"

// Write a function to check that a binary tree ↴ is a valid binary search tree
// option A: store mix maxs at each node, O(1) lookup but O(N) space
// option B: recursive func O(log N) lookup but O(1) space
// option C: incorporate this in the call stack
// umoptimized version
func checkIfBst(root *btreeNode) bool {
	// base case
	if root == nil {
		return true
	}

	// find the mins and maxs O(lgN)
	leftMax := findMax(root.left)
	rightMin := findMin(root.right)

	// check if BST property holds
	if leftMax < root.val && root.val <= rightMin {
		// recursive call to left and right substrees
		return checkIfBst(root.left) && checkIfBst(root.right)
	}
	// if above conditions not satisfied
	return false

}

// optimized based on call stack (option C)
func checkIfBst2(root *btreeNode) bool {
	// call util function
	return isBstUtil(root, math.MinInt32, math.MaxInt32)

}

// utility funcs
// recursively finds the validity tracking the bounds at each level
func isBstUtil(root *btreeNode, lowerBound, upperBound int) bool {
	// base case
	if root == nil {
		return true
	}
	// check if BST property is violated
	// Assume: all values are unique,
	// else change equality signs to <= and >=
	if root.val < lowerBound && root.val > upperBound {
		return false
	}
	// else
	return isBstUtil(root.left, lowerBound, root.val) && isBstUtil(root.right, root.val, upperBound)

}

func findMin(root *btreeNode) int {
	// for min find the leftmost node val in the left subtree
	if root.left != nil {
		return findMin(root.left)
	}
	return root.val
}

func findMax(root *btreeNode) int {
	// for max find the rightmost element in the right subtree
	if root.right != nil {
		return findMax(root.right)
	}
	return root.val
}
