package graphs

import "strings"

// LC #269

const alpha = 26

func alienOrder(words []string) string {

	// graph with adj list representation: key: vertex, val: list of edges
	adjList := map[rune][]rune{}
	// ideally this should be for unique alphabets present in dict only
	inDegree := createIndegree(words)

	// compare adjacent words and form adjacency list and in-degree array
	for i := 0; i < len(words)-1; i++ {
		r1 := []rune(words[i])
		r2 := []rune(words[i+1])
		for j := 0; j < len(r1) && j < len(r2); {
			if r1[j] == r2[j] {
				j++
			} else {
				// first char not equal
				// create edge in adj
				u := r1[j]
				v := r2[j]
				// add if v not in adj[u]
				if vertexNotInList(adjList, u, v) {
					adjList[u] = append(adjList[u], v)
					inDegree[v]++
				}
				//
				break
			}
		}
		// tricky: Check second word isn't a prefix of first word
		if strings.HasPrefix(words[i], words[i+1]) && len(r1) > len(r2) {
			return ""
		}

	}

	// now do a top sort and output the result
	queue := []rune{}
	// chose the the letters with indegree zero and add
	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	res := []rune{}
	// pop and update the inDegree
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		res = append(res, u)
		for _, v := range adjList[u] {
			// reduce inDegree
			inDegree[v]--
			if inDegree[v] == 0 {
				// add to queue
				queue = append(queue, v)
			}
		}
	}

	// check if all letters are covered are not,
	// if cycle exists there will be more alpha inDegree
	if len(res) < len(inDegree) {
		// cycle exists (num vertices > top sort vertices)
		// e.g. words = ["x", "y", "x"] i.e. all indegree > 0
		return ""
	}

	return string(res)
}

// util to create a map of the indegrees for the unique alphabets in the dict
func createIndegree(words []string) map[rune]int {
	inDegree := map[rune]int{}
	for _, w := range words {
		r := []rune(w)
		for _, v := range r {
			if _, exists := inDegree[v]; !exists {
				inDegree[v] = 0
			}
		}
	}

	return inDegree
}

// util to check if a vertex is alread in adjacency list
func vertexNotInList(adjList map[rune][]rune, u, v rune) bool {
	for _, k := range adjList[u] {
		if k == v {
			return false
		}
	}
	return true
}

// Logic this uses Kahn's algo
// impl is complicated because the vertices are not ordered nums,
// hence map is used instead of array
