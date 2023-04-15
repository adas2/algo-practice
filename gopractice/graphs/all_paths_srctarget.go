package graphs

// All possible paths in a given graph between src and target
// LC # 797

var pathList [][]int

func allPathsSourceTarget(graph [][]int) [][]int {

	// graph is given in the form of adj list
	src, target := 0, len(graph)-1
	currPath := []int{src}
	pathList = make([][]int, 0)

	dfsTargetUtil(graph, src, target, currPath)

	return pathList
}

func dfsTargetUtil(graph [][]int, src, target int, currPath []int) {
	// terminate case:
	if src == target {
		// tricky: copy into a temp slice
		path := make([]int, len(currPath))
		copy(path, currPath)
		// (otherwise 2d slice will store ref to currPath which changes over ietrations)
		pathList = append(pathList, path)
		return
	}

	// recurse for all adjacent nodes of src
	for _, v := range graph[src] {
		currPath = append(currPath, v)
		dfsTargetUtil(graph, v, target, currPath)
		// backtrack
		currPath = currPath[:len(currPath)-1]
	}
}

// logic: (similar to recursive backtracking)
// start DFS traversal from src node
// note in regular target+paht_nodes will be marked as visited first time,
// but if we backtrack we can unmark the visited node as unvisited,
// if we hit the target in a recursion we add the path traced path to list,
// since there are no cycles, the search is guaranteed to be terminate
