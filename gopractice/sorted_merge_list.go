package practice

import (
	"container/heap"
	"container/list"
	"fmt"
)

// SortedMergeList : Given K sorted lists, merge them
// return sorted list
func SortedMergeList(lists []*list.List, k int) (*list.List, error) {
	// create priority queue
	pq := make(PriorityQueue, k)

	// iterate on the lists
	for i := range lists {
		e := lists[i].Front()
		// add first element to priority queue
		pq[i] = &Item{
			value: e.Value.(int),
			index: i,
		}
		// remove the added element
		lists[i].Remove(e)

	}
	heap.Init(&pq)

	// output result list
	result := list.New()

	// ieterate till list is empty
	for pq.Len() > 0 {
		// pick the smallest element and push to result
		item := heap.Pop(&pq).(*Item)
		fmt.Println("Popped:", item.value, item.index)
		result.PushBack(item.value)

		// find the next element from same list in the pq
		if lists[item.index].Len() > 0 {
			elem := lists[item.index].Front()
			newItem := &Item{
				value: elem.Value.(int),
				index: item.index,
			}
			// push it's value/index to the list
			heap.Push(&pq, newItem)
			// and delete element pushed
			lists[item.index].Remove(elem)

		}
	}

	return result, nil
}
