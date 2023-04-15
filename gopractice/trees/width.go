package trees

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// BFS traversal queue
	nodeQueue := []*TreeNode{root}
	// stores the index of each level: 0 for root node
	levelIndexQueue := []int{0}

	maxWidth := 0

	// while nodeQueue not empty
	for len(nodeQueue) != 0 {
		levelCnt := len(nodeQueue)
		start := levelIndexQueue[0]
		end := levelIndexQueue[levelCnt-1]

		// pop nodes for this level
		for n := levelCnt; n > 0; n-- {
			node := nodeQueue[0]
			nodeQueue = nodeQueue[1:]

			i := levelIndexQueue[0]
			levelIndexQueue = levelIndexQueue[1:]

			// add children, index to nodeQueue, indexQueue
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
				levelIndexQueue = append(levelIndexQueue, 2*i)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
				levelIndexQueue = append(levelIndexQueue, (2*i)+1)
			}
		}

		// check width
		if maxWidth < (end - start + 1) {
			maxWidth = (end - start + 1)
		}

	}
	return maxWidth
}

// similar to level order traversal:
// use bfs, at each level store the nodes in a slice
// use an additional nodeQueue to store the index/count for each level
// trim the null values at the end of the slice
// record the longest slice across all levels
