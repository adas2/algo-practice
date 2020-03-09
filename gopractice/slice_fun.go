package practice

import "fmt"

type MySlice struct {
	queue  []int
	maxcap int
}

// create empty slice of type T
func createSlice(max int) *MySlice {

	s := &MySlice{
		queue:  make([]int, 0),
		maxcap: max,
	}
	// return MySlice(s)
	return s
}

// add new element at the end of slice
func (s *MySlice) addLast(elem int) {

	if s.exceedCap() {
		// capacity reached
		fmt.Printf("Capacity %d reached; try deleting\n", s.maxcap)
		return
	}

	s.queue = append(s.queue, elem)
	// fmt.Println(s)
}

func (s *MySlice) addFront(elem int) {

	if s.exceedCap() {
		// capacity reached
		fmt.Printf("Capacity %d reached; try deleting\n", s.maxcap)
		return
	}
	// v :=
	s.queue = append([]int{elem}, s.queue...)
}

// delete last element
func (s *MySlice) deleteLast() {

	if s.isEmpty() {
		fmt.Printf("Cannot delete empty queue\n")
		return
	}
	// free memory
	//  s[len(s)-1] = nil
	s.queue = s.queue[:len(s.queue)-1]
	// fmt.Println(s)
}

func (s *MySlice) deleteFront() {

	if s.isEmpty() {
		fmt.Printf("Cannot delete empty queue\n")
		return
	}
	// free memory if struct
	// s.queue[0] = nil
	s.queue = s.queue[1:]
}

func (s *MySlice) printSlice() {
	fmt.Println(s.queue)
}

func (s *MySlice) exceedCap() bool {
	// if max cap reached
	if len(s.queue) == s.maxcap {
		return true
	}
	return false
}

func (s *MySlice) isEmpty() bool {
	if len(s.queue) < 1 {
		return true
	}

	return false
}
