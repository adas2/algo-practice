package trees

// find if two trees are flip equivalent
// lc # 951

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {

	// base/termination case
	if root1 == nil && root2 == nil {
		return true
	}

	// if any one is nil, other is not
	if root1 == nil || root2 == nil {
		return false
	}

	// root val must match
	if root1.Val != root2.Val {
		return false
	}

	// case 1: no flip at root node
	// if left and right match we can ensure no flip
	// in which case we recurse on left and right subtrees
	no_flip_case := flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)

	// case 2: there was actually a flip at node
	flip_case := flipEquiv(root1.Right, root2.Left) && flipEquiv(root1.Left, root2.Right)

	// to be true at least one has to be true
	return no_flip_case || flip_case
}

// logic:
// at each recursion root node may or may not have a flip
// check root1.val == root2.val
// if no flip recurse (root1.left, root2.left) && (root1.right,root2.right)
// else if flip (root1.right, root1.left) && (root1.left, root2.right)
