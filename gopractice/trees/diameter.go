package trees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	_, dia := diameterUtil(root)
	return dia
}

// util: takes root, left longest path and right longest path and
// 	returns current height and longest diameter
func diameterUtil(root *TreeNode) (int, int) {
	// note: nil node height -1 to distinguish for leaf node (else height calc will be off by 1)
	if root == nil {
		return -1, 0
	}
	//  leaf node
	if root.Left == nil && root.Right == nil {
		return 0, 0
	}

	lheight, lpath := diameterUtil(root.Left)
	rheight, rpath := diameterUtil(root.Right)
	maxPath := Max3(lpath, rpath, lheight+rheight+2)
	currHeight := Max(rheight, lheight) + 1

	return currHeight, maxPath
}

func Max3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
	}
	if b > c {
		return b
	}
	return c
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// return max(max_longest_left, max_longest_right, path_via_root)
// path_via_root = left_height+right_height
// for leaf nodes --> return 0, height 0
// for non_leaf make height++
