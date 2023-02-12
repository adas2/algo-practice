package main

import "fmt"

// Define iterator func with properties:
//  	HasNExt() --> Bool,
// 		Next() --> interace{}

// SliceIterator
//
// RoundRobinSliceIterator
// e.g.
// a := [a1, a2, a3]
// b := [b1, b2]
// c := [c1, c2, c3]
// hasnext for c3 -> false,  all previous calls  --> true
// next order: a1 b1 c1  a2 b2 c2 a3 c3

type Iterator interface {
	HasNext() bool
	Next() interface{}

	// helper funcs to print
	fmt.Stringer
	// Serialize() string
}

// SliceIterator: simple iterator slice
type SliceIterator struct {
	container []string
}

// Methods for SliceIterator
func (s SliceIterator) HasNext() bool {
	return len(s.container) > 0
}

func (s *SliceIterator) Next() interface{} {
	temp := s.container[0]
	s.container = s.container[1:]
	return temp
}

func (s SliceIterator) String() string {
	return fmt.Sprintf("%v", s.container)
}

func NewSliceIterator(s []string) *SliceIterator {
	return &SliceIterator{
		container: s,
	}
}

// RoundRobinSliceIterator :
type RoundRobinSliceIterator struct {
	containerList []SliceIterator
	idx           int
}

// Init RoundRobinSliceIterator
func NewRoundRobinSliceIterator(strList [][]string) *RoundRobinSliceIterator {
	cList := make([]SliceIterator, len(strList))
	for i := range strList {
		cList[i] = *NewSliceIterator(strList[i])
	}

	return &RoundRobinSliceIterator{
		containerList: cList,
		idx:           0,
	}
}

func (r RoundRobinSliceIterator) HasNext() bool {

	for i := range r.containerList {
		if r.containerList[i].HasNext() {
			return true
		}
	}
	// if none has elements
	return false
}

func (r *RoundRobinSliceIterator) Next() interface{} {
	for !r.containerList[r.idx].HasNext() {
		r.idx = (r.idx + 1) % len(r.containerList)
	}
	temp := r.containerList[r.idx].Next()
	r.idx = (r.idx + 1) % len(r.containerList)

	return temp
}

func (r RoundRobinSliceIterator) String() string {
	var res string
	for i := range r.containerList {
		res += fmt.Sprintf("Index %d: %v\n", i, r.containerList[i])
	}
	return res
}

func (r RoundRobinSliceIterator) Serialize() string {
	res := "[ "
	for i := range r.containerList {
		res += r.containerList[i].String()
		res += " "
	}
	// res += fmt.Sprintf("Index %d: %v\n", i, r.containerList[i])
	return res + "]"
}

type NewIterator interface {
	Iterator
	Serialize() string
}

// test the iterators
func IteratorFun(i Iterator) {
	fmt.Println(i)
	for i.HasNext() {
		fmt.Println(i.Next())
	}
}

func IteratorFun2(n NewIterator) {
	fmt.Printf("Serialized: %s\n", n.Serialize())
}
