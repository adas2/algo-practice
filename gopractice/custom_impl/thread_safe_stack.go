package custom

import (
	"errors"
	"sync"
)

// Stack is thread-safe
type Stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s    []int
}

// NewStack create new thread safe stack
func NewStack() *Stack {
	return &Stack{sync.Mutex{}, make([]int, 0)}
}

// Push to stack
func (s *Stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

// Pop from stack
func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}
