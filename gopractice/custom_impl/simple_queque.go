package custom

import "fmt"

// Queue struct used in BST
type Queue []int

// Push element
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop element in FIFO
func (q *Queue) Pop() int {
	result := (*q)[0]
	(*q)[0] = -1 // optional
	(*q) = (*q)[1:]
	return result
}

// IsEmpty tells i queue is empty
func (q *Queue) IsEmpty() bool {
	return (len(*q) == 0)
}

// PrintQueue prints queue
func (q *Queue) PrintQueue() {
	fmt.Println(*q)
}
