package trees

import "fmt"

// LC # 545 MEDIUM

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  prints the perimeter
func boundaryOfBinaryTree(root *TreeNode) []int {
	// case 1: root nil
	if root == nil {
		return nil
	}
	res := make([]int, 1)
	res[0] = root.Val
	// print root
	fmt.Printf("%d ", root.Val)

	// case 2: if both child present
	left := root.Left
	right := root.Right
	if left != nil && right != nil {
		printLeftPeriphery(left, &res)
		printLeafNodes(left, &res)
		printLeafNodes(right, &res)
		printRightPeriphery(right, &res)
	}

	return res
}

func printLeftPeriphery(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	// case when left or right is nil
	if root.Left != nil {
		*res = append(*res, root.Val)
		fmt.Printf("%d ", root.Val)
		printLeftPeriphery(root.Left, res)
	} else if root.Right != nil {
		*res = append(*res, root.Val)
		fmt.Printf("%d ", root.Val)
		printLeftPeriphery(root.Right, res)
	}

}

func printRightPeriphery(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Right != nil {
		printRightPeriphery(root.Right, res)
		*res = append(*res, root.Val)
		fmt.Printf("%d ", root.Val)

	} else if root.Left != nil {
		printRightPeriphery(root.Left, res)
		*res = append(*res, root.Val)
		fmt.Printf("%d ", root.Val)

	}
}

func printLeafNodes(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		fmt.Printf("%d ", root.Val)
		*res = append(*res, root.Val)
	}
	printLeafNodes(root.Left, res)
	printLeafNodes(root.Right, res)
}

// logic:
// 	    	  1
//	         / \
//	        2   3
//         /   / \
//        4   5   6
//           / \  /
//          7  8  9
//
// perimeter: 1,2,7,8,9,6,3
// order of printing
// root, leftmost_nodes(left_subtree), leaf_nodes(right_subtree(lef_subtree)), leaf_nodes(left_subtree(right_subtree)), rightmost_nodes(right_subtree)
// use recursion based on above
// if left or right subtree is null
