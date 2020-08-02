package trees

// LC # 314
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  define struct for queue elements
type tuple struct {
	node *TreeNode
	idx  int
}

func verticalOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	// define a queue with <colIndex, nodeVal> as entry
	queue := make([]tuple, 0)
	// add root index, val to the queue
	// queue = append(queue, []int{0, root.Val})
	queue = append(queue, tuple{root, 0})

	// define output map to contain result ordering
	// with key = colIdx and val = {nodeVals}
	vMap := map[int][]int{}

	minCol, maxCol := 0, 0
	// do a BFS now
	for len(queue) != 0 {
		// dequeue
		colIdx := queue[0].idx
		currNode := queue[0].node
		queue = queue[1:]

		// update the map
		if _, exists := vMap[colIdx]; exists {
			vMap[colIdx] = append(vMap[colIdx], currNode.Val)
		} else {
			vMap[colIdx] = []int{currNode.Val}
		}

		// update min,max indx
		if minCol > colIdx {
			minCol = colIdx
		}
		if maxCol < colIdx {
			maxCol = colIdx
		}

		// push the children of the node popped
		if currNode.Left != nil {
			queue = append(queue, tuple{currNode.Left, colIdx - 1})
		}
		if currNode.Right != nil {
			queue = append(queue, tuple{currNode.Right, colIdx + 1})
		}
	}

	res := [][]int{}
	// now output the results from minCol to maxCol
	for i := minCol; i <= maxCol; i++ {
		if _, exists := vMap[i]; exists {
			res = append(res, vMap[i])
		}
	}

	return res
}

// Logic:
// This is somewhat similar to leve-order traversal execpt the fact that the results have to sorted
// in columns index order, i.e. if root index = 0, left index =-1 right index =+1
// BFS approach with extra information stored in the queue about the column index
// Maintain a map with column index as key and list of nodes as value,
// because of BFS property, nodes higher level will be pushed first.
// to display the results from min_column index to max_column index keeo track of them
// the last step avoids sorting of the map elements
