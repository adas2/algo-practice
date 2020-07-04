package graphs

import (
	custom "adas2.io/practice/custom_impl"
)

// use Kahn's algo by using indegree of virtices
func topSortGraph(g *Graph, inDegree []int) []int {

	// use a queue to store all virtices with in-dgree zero
	q := custom.Queue{}
	for i := range inDegree {
		if inDegree[i] == 0 {
			q.Push(i)
		}
	}
	res := []int{}
	// while queue is not empty
	for q.IsEmpty() == false {
		// pop a virtex and add to result
		u := q.Pop()
		res = append(res, u)
		// decrement the in-degrees of all it adjacent nodes
		for _, v := range g.adjV[u] {
			inDegree[v]--
			// add adj nodes with in-degree zero to this queue
			if inDegree[v] == 0 {
				q.Push(v)
			}
		}
	}

	return res

}

// time complexity: O(V+E) since for every virtex its edges is visited once
