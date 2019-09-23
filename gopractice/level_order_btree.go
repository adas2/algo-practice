package practice

import (
	"fmt"
	"io"
)

// level order traversal for a Binary Tree
// BFS style traversal
// Time: O(N) Space: O(N)
func PrintLevelOrder(w io.Writer, tree *bTree) {
	if tree.root == nil {
		return
	}
	// implement a queue and store the nodes
	// hint: like a BFS traversal
	queue := make([]*btreeNode, 0)
	// insert root
	queue = append(queue, tree.root)

	// while queue not empty
	for len(queue) != 0 {
		// piick front node
		node := queue[0]

		// consume/output node
		fmt.Fprintf(w, "%d ", node.val)

		// dequeue when done with this node [CHECK?]
		queue = queue[1:]

		// add children to queue
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

// Queue methods
// Enqueue: queue = append(queue, value)
// Dequeue: queue = queue[1:]
// Peek: queue[0]
// IsEmpty: if len(queue) == 0 return true
