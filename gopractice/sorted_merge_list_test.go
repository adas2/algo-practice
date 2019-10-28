package practice

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedMergeList(t *testing.T) {
	// initialize lists
	lists := initLists()
	fmt.Println("The Input lists:")
	for i := range lists {
		for e := lists[i].Front(); e != nil; e = e.Next() {
			// do something with e.Value
			fmt.Printf("%v ", e.Value)
		}
		fmt.Println()

	}
	result, err := SortedMergeList(lists, 2)
	assert.Nil(t, err)
	for e := result.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()

}

func initLists() []*list.List {
	lists := []*list.List{}
	l1 := list.New()
	l2 := list.New()
	// l1.Init()
	// l2.Init()
	l1.PushBack(3)
	l1.PushBack(6)
	l1.PushBack(9)

	l2.PushBack(1)
	l2.PushBack(5)
	l2.PushBack(10)

	// l1.Init()
	lists = append(lists, l1, l2)

	return lists
}
