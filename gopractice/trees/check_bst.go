package trees

import (
	"fmt"
	"math"
)

// Write a function to check if a binary tree is a valid BST
// option A: store min maxs at each node, O(1) lookup but O(N) space (augment tree needed)
// option B: recursive func O(log N) lookup but O(1) space
// option C: incorporate this in the call stack

// umoptimized version (option B)
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
	return isBstUtil(root, math.MinInt64, math.MaxInt64)

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
	if root.val > lowerBound && root.val < upperBound {

		fmt.Println("Root: ", root.val, upperBound, lowerBound)
		return isBstUtil(root.left, lowerBound, root.val) && isBstUtil(root.right, root.val, upperBound)
	}

	// else
	return false
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
