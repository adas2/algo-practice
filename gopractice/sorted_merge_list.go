package practice

import (
	"container/heap"
	"container/list"
	"fmt"

	custom "adas2.io/practice/custom_impl"
)

// SortedMergeList: Given K sorted lists, merge them
// return sorted list
func SortedMergeList(lists []*list.List, k int) (*list.List, error) {
	// create priority queue
	pq := make(custom.PriorityQueue, k)

	// iterate on the lists
	for i := range lists {
		e := lists[i].Front()
		// add first element to priority queue
		pq[i] = &custom.Item{
			Value: e.Value.(int),
			Index: i,
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
		item := heap.Pop(&pq).(*custom.Item)
		fmt.Println("Popped:", item.Value, item.Index)
		result.PushBack(item.Value)

		// find the next element from same list in the pq
		if lists[item.Index].Len() > 0 {
			elem := lists[item.Index].Front()
			newItem := &custom.Item{
				Value: elem.Value.(int),
				Index: item.Index,
			}
			// push it's value/index to the list
			heap.Push(&pq, newItem)
			// and delete element pushed
			lists[item.Index].Remove(elem)

		}
	}

	return result, nil
}
