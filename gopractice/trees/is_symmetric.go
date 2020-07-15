package trees

// Is BT symmetric?

// recrusive O(N)
func isSymmetric(r1 *btreeNode, r2 *btreeNode) bool {

	if r1 == nil && r2 == nil {
		return true
	} else if r1 == nil || r2 == nil {
		return false
	}
	// non-nil r1 and r2 check
	if r1.val != r2.val {
		return false
	}

	return isSymmetric(r1.left, r2.right) && isSymmetric(r1.right, r2.left)
}

// iterative version t=O(N), s=O(N)
// based on level order traversal
func isSymmetricIterative(r *btreeNode) bool {
	// define a queue
	q := []*btreeNode{}
	// push left anf right child values and iterated
	if r.left != nil {
		// push
		q = append(q, r.left)
	}
	if r.right != nil {
		// push
		q = append(q, r.right)
	}
	var r1, r2 *btreeNode
	for len(q) > 0 {
		if len(q) < 2 { // no enought nodes
			break
		}
		// pop first 2
		r1 = q[0]
		r2 = q[1]
		q = q[2:]
		if r1 == nil && r2 == nil {
			continue
		}
		if r1 == nil || r2 == nil {
			return false
		}
		if r1.val != r2.val {
			return false
		}
		// push in the order lrrl
		q = append(q, r1.left, r2.right, r1.right, r2.left)
	}
	// if queue empty all nodes processed
	if len(q) == 0 {
		return true
	}
	return false
}
